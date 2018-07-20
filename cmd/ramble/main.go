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
	app.Action = chat
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "n, name",
			Usage: "unique user name",
		},
		cli.StringFlag{
			Name:  "a, addr",
			Usage: "address of remote to connect to",
		},
		cli.UintFlag{
			Name:  "p, port",
			Usage: "port to listen for messages on",
			Value: 3264,
		},
	}

	app.Run(os.Args)
}

func chat(c *cli.Context) error {
	ramble := ramble.New(c.String("name"))

	if err := ramble.Listen(fmt.Sprintf(":%d", c.Uint("port"))); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	if err := ramble.Connect(c.String("addr")); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
