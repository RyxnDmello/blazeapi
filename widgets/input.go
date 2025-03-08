package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type input struct {
	label        string
	placeholder  string
	handleChange func(text string)
	handleAccept func(text string, lastChar rune) bool
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

func NewInput() *input {
	return &input{
		label:        "",
		placeholder:  "",
		handleAccept: nil,
		handleInput:  nil,
	}
}

func (input *input) SetLabel(label string) *input {
	input.label = label
	return input
}

func (input *input) SetPlaceholder(placeholder string) *input {
	input.placeholder = placeholder
	return input
}

func (input *input) HandleChange(handleChange func(text string)) *input {
	input.handleChange = handleChange
	return input
}

func (input *input) HandleAcceptance(handleAccept func(text string, lastChar rune) bool) *input {
	input.handleAccept = handleAccept
	return input
}

func (input *input) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *input {
	input.handleInput = handleInput
	return input
}

func (input *input) Render() (field *tview.InputField) {
	field = tview.
		NewInputField().
		SetLabel(input.label).
		SetPlaceholder(input.placeholder).
		SetChangedFunc(input.handleChange).
		SetAcceptanceFunc(input.handleAccept).
		SetFieldStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetLabelStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		).
		SetPlaceholderStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		)

	field.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetInputCapture(input.handleInput).
		SetBackgroundColor(tcell.ColorBlack)

	return field
}
