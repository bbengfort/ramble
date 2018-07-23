package main

import (
	"fmt"
	"os"

	"github.com/bbengfort/ramble"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	// Load the .env file if it exists
	godotenv.Load()

	// Instantiate the command line application
	app := cli.NewApp()
	app.Name = "ramble"
	app.Version = ramble.PackageVersion
	app.Usage = "a streaming gRPC point to point chat system"
	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "run the chat server",
			Action: serve,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "p, port",
					Usage: "port to listen for messages on",
					Value: 3265,
				},
				cli.UintFlag{
					Name:  "verbosity",
					Usage: "verbosity of log statements, lower is more verbose",
					Value: 2,
				},
			},
		},
		{
			Name:   "chat",
			Usage:  "run the chat client program",
			Action: chat,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "n, name",
					Usage: "unique user name",
				},
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address of remote to connect to",
				},
			},
		},
		{
			Name:   "bench",
			Usage:  "run a benchmark against the chat server",
			Action: bench,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address of remote to connect to",
				},
				cli.IntFlag{
					Name:  "c, clients",
					Usage: "number of concurrent clients",
					Value: 4,
				},
				cli.IntFlag{
					Name:  "m, messages",
					Usage: "messages per client to send to server",
					Value: 5000,
				},
			},
		},
	}

	app.Run(os.Args)
}

//===========================================================================
// Actions
//===========================================================================

func serve(c *cli.Context) error {
	ramble.SetLogLevel(uint8(c.Uint("verbosity")))

	server := ramble.NewServer(c.Uint("port"))
	if err := server.Listen(); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

func chat(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		name = ramble.Prompt("Enter username: ")
	}

	addr := c.String("addr")
	if addr == "" {
		addr = ramble.Prompt("Enter server address: ")
	}

	console, err := ramble.NewConsole(name, addr)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	defer console.Close()

	if err := console.Run(); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

func bench(c *cli.Context) error {
	addr := c.String("addr")
	if addr == "" {
		return cli.NewExitError("must specify an address to connect to", 1)
	}

	bench, err := ramble.NewBenchmark(c.Int("clients"), c.Int("messages"), addr)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(bench.String())
	return nil
}
