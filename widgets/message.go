package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type message struct {
	label       string
	handleInput func(event *tcell.EventKey) *tcell.EventKey
}

func NewMessage() *message {
	return &message{
		label:       "",
		handleInput: nil,
	}
}

func (message *message) SetLabel(label string) *message {
	message.label = label
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
		SetText(message.label).
		SetTextAlign(tview.AlignLeft).
		SetTextStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		)

	element.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	return element
}
