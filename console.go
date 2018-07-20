package ramble

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bbengfort/ramble/pb"
	"github.com/jroimartin/gocui"
	"google.golang.org/grpc"
)

//===========================================================================
// Console Program
//===========================================================================

// Singleton reference to console program
var console *Console

// NewConsole creates the terminal UI in 256 colors, instantiates the layout
// and widgets, and binds the keys to event handlers. Not thread safe!
func NewConsole(name, addr string) (*Console, error) {
	var err error

	// Create the console properties
	console = &Console{address: addr, username: name}
	if console.cui, err = gocui.NewGui(gocui.Output256); err != nil {
		return nil, err
	}

	// Create the console layout and edit properties
	console.cui.Cursor = true
	console.CreateLayout()

	// Bind the keys in the ramble console app
	if err := console.BindKeys(); err != nil {
		return nil, err
	}

	return console, nil
}

// Console implments a terminal UI client to the chat server.
type Console struct {
	cui      *gocui.Gui           // terminal UI
	address  string               // address of the chat server
	username string               // username of local user
	sequence int64                // the sequence id for each message
	conn     *grpc.ClientConn     // connection to the chat service
	client   pb.RambleClient      // ramble service client
	stream   pb.Ramble_ChatClient // the stream to the ramble chat
}

// Close the console program and cleanup the screen.
func (c *Console) Close() {
	c.stream.CloseSend()
	c.conn.Close()
	c.cui.Close()
}

// Run the console program's main loop and return any errors.
func (c *Console) Run() error {
	if err := c.connect(); err != nil {
		return err
	}

	if err := c.cui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

// Connect to the chat server and create the chat stream
func (c *Console) connect() (err error) {
	if c.conn, err = grpc.Dial(c.address, grpc.WithInsecure()); err != nil {
		return err
	}

	c.client = pb.NewRambleClient(c.conn)
	if c.stream, err = c.client.Chat(context.Background()); err != nil {
		return err
	}

	// Run the listener function
	go c.listen()
	return nil
}

// Listen for messages from the chat server and write them to the messages box.
func (c *Console) listen() {
	for {
		in, err := c.stream.Recv()
		if err == io.EOF {
			// No more messages from the server are coming
			in = &pb.ChatMessage{
				Sender:    "system",
				Timestamp: ChatTime(),
				Message:   Colorize(160, fmt.Sprintf("disconnected from %s", c.address)),
			}
		}
		if err != nil {
			// Something bad happened, panic!
			c.Close()
			log.Fatalf("could not receive chat message: %s", err)
		}

		if err := messages.Append(in); err != nil {
			c.Close()
			log.Fatalf("could not receive chat message: %s", err)
		}

		c.cui.Update(func(g *gocui.Gui) error {
			return nil
		})
	}
}

//===========================================================================
// Helper Functions
//===========================================================================

// Colorize the specified text in 256 terminal colors
func Colorize(color int, text string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", color, text)
}

// Prompt for information from the command line.
func Prompt(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()

		text := scanner.Text()
		if len(text) > 0 {
			return text
		}
	}
}
