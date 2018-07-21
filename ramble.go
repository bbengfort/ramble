package ramble

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/bbengfort/ramble/pb"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PackageVersion of the Ramble app
const PackageVersion = "1.0"

// ServerName is reserved for use by system messages
const ServerName = "server"

// NewServer creates a chat server to distribute messages to all streaming
// clients that connect via gRPC.
func NewServer(port uint) *Ramble {
	ramble := &Ramble{port: port}
	ramble.clients = make(map[uint64]chan *pb.ChatMessage)

	return ramble
}

// Ramble implements the RambleService
type Ramble struct {
	sync.Mutex        // Protect concurrent access to sequence and client manager
	port       uint   // Port to listen on for client connections
	sequence   int64  // Message sequence number for oredering evaluation
	clientID   uint64 // Assign unique ids to all clients

	// All currently connected clients
	clients map[uint64]chan *pb.ChatMessage
}

// Listen for chat service messages
func (r *Ramble) Listen() error {
	addr := fmt.Sprintf(":%d", r.port)
	sock, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("could not listen on %s", addr)
	}
	defer sock.Close()

	srv := grpc.NewServer()
	pb.RegisterRambleServer(srv, r)
	return srv.Serve(sock)
}

// Chat handles chat stream clients.
func (r *Ramble) Chat(stream pb.Ramble_ChatServer) error {
	clientID, messages := r.connectClient()

	// Run a go routine that sends chat messages for other clients
	go func() {
		for msg := range messages {
			if err := stream.Send(msg); err != nil {
				// Note that we don't really care about send errors
				log.Printf("could not send msg %d to client %d: %s", msg.Sequence, clientID, err)
			}
		}
	}()

	for {
		in, err := stream.Recv()
		if err != nil {
			// On any receiving errors, disconnect the client
			r.disconnectClient(clientID)

			if err == io.EOF {
				// This means that the client has disconnected gracefully
				return nil
			}

			// Log the error and return it
			log.Printf("error receiving from client %d: %s", clientID, err)
			return err
		}

		// broadcast message to all connected clients
		r.broadcast(in)
	}
}

// Ping handles ping stream clients
func (r *Ramble) Ping(stream pb.Ramble_PingServer) error {
	fmt.Println("hidy ping!")
	return nil
}

// broadcast a message, incrementing the sequence number and sending to all
// currently connected clients (including the sender of the message).
func (r *Ramble) broadcast(msg *pb.ChatMessage) {
	r.Lock()
	defer r.Unlock()

	// TODO: create a synchronized sequence data structure
	r.sequence++
	msg.Sequence = r.sequence

	// TODO: This should be a readlock only
	for _, send := range r.clients {
		send <- msg
	}

	log.Printf("msg %d from %s: %s", msg.Sequence, msg.Sender, msg.Message)
}

// connectClient creates a unique ID and associates it with a message channel
// then sends a connection broadcast message to all clients.
func (r *Ramble) connectClient() (uint64, chan *pb.ChatMessage) {
	r.Lock()
	defer r.Unlock()

	r.clientID++
	r.clients[r.clientID] = make(chan *pb.ChatMessage, 10)

	// TODO: Create a synchronized sequence data structure
	r.sequence++
	msg := &pb.ChatMessage{
		Sequence:  r.sequence,
		Sender:    ServerName,
		Timestamp: ChatTime(),
		Message:   fmt.Sprintf("client %d has connected", r.clientID),
	}

	// TODO: This should be a read lock only
	for _, send := range r.clients {
		send <- msg
	}

	log.Printf("client %d connected\n", r.clientID)
	return r.clientID, r.clients[r.clientID]
}

// disconnectClient removes the messages channel from the client list and
// broadcasts a disconnect message to all clients.
func (r *Ramble) disconnectClient(clientID uint64) {
	r.Lock()
	defer r.Unlock()

	close(r.clients[clientID])
	delete(r.clients, clientID)

	// TODO: Create a synchronized sequence data structure
	r.sequence++
	msg := &pb.ChatMessage{
		Sequence:  r.sequence,
		Sender:    ServerName,
		Timestamp: ChatTime(),
		Message:   fmt.Sprintf("client %d has disconnected", clientID),
	}

	// TODO: This should be a read lock only
	for _, send := range r.clients {
		send <- msg
	}

	log.Printf("client %d disconnected\n", clientID)
}
