package main

import (
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
					Value: 3264,
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
	}

	app.Run(os.Args)
}

//===========================================================================
// Actions
//===========================================================================

func serve(c *cli.Context) error {
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
