package message

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Message struct {
	text            string
	align           int
	wrap            bool
	wordWrap        bool
	design          Design
	handleChange    func()
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewMessage() *Message {
	return &Message{
		text:            "",
		align:           0,
		wrap:            true,
		wordWrap:        true,
		design:          NewDesign(),
		handleChange:    nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (message *Message) SetText(text string) *Message {
	message.text = text
	return message
}

func (message *Message) SetAlignment(alignment int) *Message {
	message.align = alignment
	return message
}

func (message *Message) SetWrap(wrap bool) *Message {
	message.wrap = wrap
	return message
}

func (message *Message) SetWordWrap(wordWrap bool) *Message {
	message.wordWrap = wordWrap
	return message
}

func (message *Message) SetDesign(design Design) *Message {
	message.design = design
	return message
}

func (message *Message) HandleChange(handleChange func()) *Message {
	message.handleChange = handleChange
	return message
}

func (message *Message) HandleSwitch(handleSwitch func()) *Message {
	message.handleSwitch = handleSwitch
	return message
}

func (message *Message) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Message {
	message.handleInput = handleInput
	return message
}

func (message *Message) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Message {
	message.handleDimension = handleDimension
	return message
}

func (message *Message) View() *tview.TextView {
	widget := tview.
		NewTextView().
		SetText(message.text).
		SetWrap(message.wrap).
		SetDynamicColors(true).
		SetTextAlign(message.align).
		SetWordWrap(message.wordWrap).
		SetTextStyle(message.design.GetStyle())

	widget.
		SetChangedFunc(func() {
			if message.handleChange != nil {
				message.handleChange()
			}
		}).
		SetDoneFunc(func(key tcell.Key) {
			if message.handleSwitch != nil && key == tcell.KeyTAB {
				message.handleSwitch()
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(message.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if message.handleSwitch != nil && event.Key() == tcell.KeyTab {
				message.handleSwitch()
			}

			if message.handleInput != nil {
				return message.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range message.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), false)
			}

			if message.handleDimension == nil {
				return x, y, width, height
			}

			return message.handleDimension(screen, x, y, width, height)
		})

	return widget
}
