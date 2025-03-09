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

func NewMessage() *message {
	return &message{
		text:        "",
		alignment:   tview.AlignLeft,
		handleInput: nil,
	}
}

func (message *message) SetText(text string) *message {
	message.text = text
	return message
}

func (message *message) SetAlignment(alignment int) *message {
	message.alignment = alignment
	return message
}

func (message *message) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *message {
	message.handleInput = handleInput
	return message
}

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
