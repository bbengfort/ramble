package ramble

import (
	"fmt"
	"io"
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
const PackageVersion = "0.2"

// NewServer creates a chat server to distribute messages to all streaming
// clients that connect via gRPC.
func NewServer(port uint) *Ramble {
	ramble := &Ramble{port: port}
	ramble.clients = make(map[string]chan *pb.ChatMessage)

	return ramble
}

// Ramble implements the RambleService
type Ramble struct {
	sync.Mutex
	port    uint
	clients map[string]chan *pb.ChatMessage
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
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// TODO: send message to everyone else!
		fmt.Println(in)

		ack := &pb.ChatMessage{
			Sender:    "system",
			Timestamp: ChatTime(),
			Message:   "message received",
		}
		if err := stream.Send(ack); err != nil {
			return err
		}
	}
}

// Ping handles ping stream clients
func (r *Ramble) Ping(stream pb.Ramble_PingServer) error {
	fmt.Println("hidy ping!")
	return nil
}

// ChatTime returns the current timestamp formatted for the chat window.
func ChatTime() string {
	return time.Now().Format("15:04:05")
}
