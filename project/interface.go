package project

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Input(placeholder string, handleAcceptance func(textToCheck string, lastChar rune) bool, handleInput func(event *tcell.EventKey) *tcell.EventKey) (projectInput *tview.InputField) {
	projectInput = tview.
		NewInputField().
		SetPlaceholder(placeholder).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetAcceptanceFunc(handleAcceptance)

	projectInput.
		SetBorder(true).
		SetInputCapture(handleInput).
		SetBorderPadding(0, 0, 1, 1)

	return projectInput
}

func Button(label string, handleSelect func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (createButton *tview.Button) {
	createButton = tview.
		NewButton(fmt.Sprintf(" %s ", label)).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetSelectedFunc(handleSelect)

	createButton.
		SetBorder(true).
		SetInputCapture(handleInput)

	return createButton
}
