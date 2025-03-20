package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type message struct {
	text        string
	alignment   int
	handleInput func(event *tcell.EventKey) *tcell.EventKey
}

// NewMessage creates and returns an instance of the message type.
//
// The message widget is used to display text content in Blaze. Properties
// such as the text, alignment, and input handler can be customized. The widget
// aligns with the features and requirements of Blaze. Under the hood, it uses
// tview.TextView to render the user interface.
//
// # Properties
//
//	text        string
//	alignment   int
//	handleInput func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*message
//
// # Usage
//
//	message := widgets.NewMessage()
func NewMessage() *message {
	return &message{
		text:        "",
		alignment:   tview.AlignLeft,
		handleInput: nil,
	}
}

// SetText sets the text content of the message type.
//
// The text content can be set or updated dynamically with this function.
// Method chaining is supported.
//
// # Parameters
//
//	text string
//
// # Returns
//
//	*message
//
// # Usage
//
//	message := widgets.NewMessage().SetText("")
func (message *message) SetText(text string) *message {
	message.text = text
	return message
}

// SetAlignment sets the text alignment for the message type.
//
// The alignment of the text can be set to control its positioning. Method
// chaining is supported.
//
// # Parameters
//
//	alignment int
//
// # Returns
//
//	*message
//
// # Usage
//
//	message := widgets.NewMessage().SetAlignment(tview.AlignLeft)
func (message *message) SetAlignment(alignment int) *message {
	message.alignment = alignment
	return message
}

// HandleInput assigns a custom function to process keyboard events.
//
// Keyboard behavior can be customized with a tcell EventKey handler. If nil,
// tview defaults are used. Method chaining is supported.
//
// # Parameters
//
//	handleInput func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*message
//
// # Usage
//
//	message := widgets.
//	     NewMessage().
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     )
func (message *message) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *message {
	message.handleInput = handleInput
	return message
}

// Render creates and returns a tview.TextView with the configured properties.
//
// A tview.TextView is built with the set properties. This function renders
// the message for display in Blaze.
//
// # Returns
//
// *tview.TextView
//
// # Usage
//
//	message := widgets.
//	     NewMessage().
//	     SetText("").
//	     SetAlignment(tview.AlignLeft).
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     ).
//	     Render()
func (message *message) Render() *tview.TextView {
	element := tview.
		NewTextView().
		SetWrap(true).
		SetWordWrap(true).
		ScrollToBeginning().
		SetDynamicColors(true).
		SetText(message.text).
		SetTextAlign(message.alignment).
		SetTextStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		)

	element.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetInputCapture(message.handleInput)

	return element
}
