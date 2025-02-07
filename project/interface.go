package project

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Input(placeholder string, handleAcceptance func(textToCheck string, lastChar rune) bool, handleInput func(event *tcell.EventKey) *tcell.EventKey) (input *tview.InputField) {
	input = tview.
		NewInputField().
		SetPlaceholder(placeholder).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetAcceptanceFunc(handleAcceptance)

	input.
		SetBorder(true).
		SetInputCapture(handleInput).
		SetBorderPadding(0, 0, 1, 1)

	return input
}

func Button(label string, handleSelect func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (button *tview.Button) {
	button = tview.
		NewButton(fmt.Sprintf(" %s ", label)).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetSelectedFunc(handleSelect)

	button.
		SetBorder(true).
		SetInputCapture(handleInput)

	return button
}
