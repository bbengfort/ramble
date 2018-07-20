package ramble

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/bbengfort/ramble/pb"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PackageVersion of the Ramble app
const PackageVersion = "0.1"

// New ramble service
func New(name string) *Ramble {
	return &Ramble{Name: name}
}

// Ramble implements the RambleService
type Ramble struct {
	Name string // name of the local client
}

// PingService sends a message on a routine tick
func (r *Ramble) PingService(addr string, delay time.Duration) error {
	var sequence int64
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect to %s: %s", addr, err)
	}

	client := pb.NewRambleClient(conn)
	stream, err := client.Ping(context.Background())
	if err != nil {
		return err
	}

	ticker := time.NewTicker(delay)
	for {
		ts := <-ticker.C
		sequence++
		if err = stream.Send(&pb.PingRequest{
			Sequence: sequence,
			Sender:   r.Name,
			Ttl:      30,
			Payload:  []byte(ts.String()),
		}); err != nil {
			return err
		}
	}
}

// Listen for chat service messages
func (r *Ramble) Listen(addr string) error {
	sock, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("could not listen on %s", addr)
	}
	defer sock.Close()

	srv := grpc.NewServer()
	pb.RegisterRambleServer(srv, r)
	go srv.Serve(sock)
	return nil
}

// Chat handles chat stream clients
func (r *Ramble) Chat(stream pb.Ramble_ChatServer) error {
	fmt.Println("hidy ho!")
	return nil
}

// Ping handles ping stream clients
func (r *Ramble) Ping(stream pb.Ramble_PingServer) error {
	fmt.Println("hidy ping!")
	return nil
}
