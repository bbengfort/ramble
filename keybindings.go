package ramble

import (
	"github.com/jroimartin/gocui"
)

// BindKeys sets the key bindings for the application.
func (c *Console) BindKeys() error {

	// Exit on CTRL+C
	if err := c.cui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Togle View bindings
	for _, name := range []string{"", MsgsView, ChatView} {
		if err := c.cui.SetKeybinding(name, gocui.KeyTab, gocui.ModNone, toggleView); err != nil {
			return err
		}
	}

	// Send messages on enter in the chat box
	if err := c.cui.SetKeybinding(ChatView, gocui.KeyEnter, gocui.ModNone, chatbox.Send); err != nil {
		return err
	}

	// Scroll up with cursor up in messages
	if err := c.cui.SetKeybinding(MsgsView, gocui.KeyArrowUp, gocui.ModNone, messages.ScrollUp); err != nil {
		return err
	}

	// Scroll down with cursor down in messages
	if err := c.cui.SetKeybinding(MsgsView, gocui.KeyArrowDown, gocui.ModNone, messages.ScrollDown); err != nil {
		return err
	}

	return nil
}

//===========================================================================
// Global Keypress Event Handlers
//===========================================================================

// Quit the program on CTRL+C
func quit(g *gocui.Gui, v *gocui.View) error {
	console.Close()
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
