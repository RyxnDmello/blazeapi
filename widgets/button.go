package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type button struct {
	label        string
	handleSelect func()
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

func NewButton() *button {
	return &button{
		label:        "",
		handleSelect: nil,
		handleInput:  nil,
	}
}

func (button *button) SetLabel(label string) *button {
	button.label = label
	return button
}

func (button *button) HandleSelect(handleSelect func()) *button {
	button.handleSelect = handleSelect
	return button
}

func (button *button) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *button {
	button.handleInput = handleInput
	return button
}

func (button *button) Render() *tview.Button {
	element := tview.
		NewButton(button.label).
		SetSelectedFunc(button.handleSelect).
		SetStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetActivatedStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetDisabledStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		)

	element.
		SetBorder(true).
		SetInputCapture(button.handleInput).
		SetBackgroundColor(tcell.ColorBlack)

	return element
}
