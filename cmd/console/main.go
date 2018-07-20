package main

import (
	"log"

	"github.com/bbengfort/ramble"
	"github.com/jroimartin/gocui"
)

// View name constants
var (
	MsgsView = ramble.MsgsView
	ChatView = ramble.ChatView
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	ramble.NewLayout(g)

	if err := ramble.BindKeys(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
