package ramble

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bbengfort/ramble/pb"
	"github.com/jroimartin/gocui"
)

// View names for the major layout components
const (
	MsgsView    = "messages"
	ChatView    = "chatbox"
	HelpView    = "helpbox"
	HistorySize = 150
)

// Singleton widget variables
var (
	messages *Messages
	chatbox  *ChatBox
	helpbox  *HelpWidget
)

// NewLayout creates the messages and chatbox widgets then sets the layout
// manager on the GUI application.
func NewLayout(g *gocui.Gui) {
	messages = new(Messages)
	messages.name = MsgsView
	messages.lines = make([]*pb.ChatMessage, 0)
	messages.users = make(map[string]int)

	chatbox = new(ChatBox)
	chatbox.name = ChatView

	helpbox = &HelpWidget{text: "TAB: toggle view | CTRL+C: exit"}
	helpbox.name = HelpView

	g.SetManager(messages, chatbox, helpbox)
}

//===========================================================================
// Widget Implementation
//===========================================================================

// Widget implements a simple GoCUI box widget
type Widget struct {
	name string
	x, y int
	w, h int
}

// Layout draws the widget on the screen
func (w *Widget) Layout(g *gocui.Gui) error {
	_, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	return nil
}

// View returns the view associated with the widget
func (w *Widget) View(g *gocui.Gui) (*gocui.View, error) {
	return g.View(w.name)
}

//===========================================================================
// Messages Widget
//===========================================================================

// Messages implements the message reader interface
type Messages struct {
	Widget
	lines []*pb.ChatMessage
	users map[string]int
}

// Layout draws a full width and tall height box, leaving room only for the
// chatbox and helpbox widgets at the bottom of the screen.
func (w *Messages) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	w.w = maxX - 1
	w.h = maxY - 8

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Wrap = true
	v.Autoscroll = true
	v.Clear()

	// Find maximum username length
	maxLen := 0
	for name := range w.users {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	// Create the message format string
	msgFmt := fmt.Sprintf("%%s  %%-%ds %%s", maxLen)

	for _, msg := range w.lines {
		name := Colorize(w.users[msg.Sender], "@"+msg.Sender)
		line := fmt.Sprintf(msgFmt, msg.Timestamp, name, msg.Message)
		fmt.Fprintln(v, line)
	}
	return nil
}

// Append a message to the messages window, limiting the history size.
func (w *Messages) Append(msg *pb.ChatMessage) (err error) {
	// If user is not in users map, assign a random color
	if _, ok := w.users[msg.Sender]; !ok {
		w.users[msg.Sender] = rand.Intn(228) + 1
	}

	// Append messages to lines
	w.lines = append(w.lines, msg)

	// Truncate the number of lines if greater than the history
	if len(w.lines) > HistorySize {
		idx := len(w.lines) - HistorySize
		w.lines = w.lines[idx:]
	}

	// Return no errors
	return nil
}

// ScrollUp scrolls the messages up by one line
func (w *Messages) ScrollUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// ScrollDown scrolls the messages down by one line
func (w *Messages) ScrollDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

//===========================================================================
// ChatBox Widget
//===========================================================================

// ChatBox implements the editable box to write messages in
type ChatBox struct {
	Widget
}

// Layout draws a full width box 4 units high at the bottom of the screen
func (w *ChatBox) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	w.y = maxY - 7
	w.w = maxX - 1
	w.h = 4

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Editable = true
	return nil
}

// Clear the ChatBox and reset the cursor.
func (w *ChatBox) Clear(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	v.SetCursor(0, 0)
	return nil
}

// Send the message currently in the ChatBox (usually bound to Enter)
func (w *ChatBox) Send(g *gocui.Gui, v *gocui.View) (err error) {
	var line string

	// Get the line from the chatbox
	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		line = ""
	}

	// Strip spaces from the line and return if no message is sent.
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	msg := &pb.ChatMessage{
		Sender:    "system",
		Timestamp: time.Now().Format("15:04:05"),
		Message:   line,
	}

	// Append the line to the messages view
	if err = messages.Append(msg); err != nil {
		return err
	}

	return w.Clear(g, v)
}

//===========================================================================
// Help Widget
//===========================================================================

// HelpWidget displays the help text at the bottom of the screen.
type HelpWidget struct {
	Widget
	text string
}

// Layout draws a full width box 1u high at the bottom of the screen
func (w *HelpWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	w.y = maxY - 3
	w.w = maxX - 1
	w.h = 2

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	v.Frame = false
	v.Clear()

	// Print help text in the center of the screen using crazy fmt string
	text := Colorize(247, w.text)
	tlen := len(text) + (len(text) - len(w.text))
	fmt.Fprintf(v, "%[1]*s", -w.w, fmt.Sprintf("%[1]*s", (w.w+tlen)/2, text))
	return nil
}
