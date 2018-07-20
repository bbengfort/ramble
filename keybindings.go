package ramble

import (
	"github.com/jroimartin/gocui"
)

// BindKeys sets the key bindings for the application.
func BindKeys(g *gocui.Gui) error {

	// Exit on CTRL+C
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Togle View bindings
	for _, name := range []string{"", MsgsView, ChatView} {
		if err := g.SetKeybinding(name, gocui.KeyTab, gocui.ModNone, toggleView); err != nil {
			return err
		}
	}

	// Send messages on enter in the chat box
	if err := g.SetKeybinding(ChatView, gocui.KeyEnter, gocui.ModNone, chatbox.Send); err != nil {
		return err
	}

	// Scroll up with cursor up in messages
	if err := g.SetKeybinding(MsgsView, gocui.KeyArrowUp, gocui.ModNone, messages.ScrollUp); err != nil {
		return err
	}

	// Scroll down with cursor down in messages
	if err := g.SetKeybinding(MsgsView, gocui.KeyArrowDown, gocui.ModNone, messages.ScrollDown); err != nil {
		return err
	}

	return nil
}

//===========================================================================
// Global Keypress Event Handlers
//===========================================================================

// Quit the program on CTRL+C
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Switch between messages and chatbox on TAB
func toggleView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == MsgsView {
		_, err := g.SetCurrentView(ChatView)
		return err
	}

	_, err := g.SetCurrentView(MsgsView)
	return err
}
