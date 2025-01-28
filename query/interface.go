package query

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RequestInput(handleAccept func(textToCheck string, lastChar rune) bool, handleInput func(event *tcell.EventKey) *tcell.EventKey) (requestInput *tview.InputField) {
	input := tview.
		NewInputField().
		SetPlaceholder("Enter URL").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetAcceptanceFunc(handleAccept)

	input.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetBackgroundColor(tcell.ColorBlack).
		SetInputCapture(handleInput)

	return input
}

func RequestBody() (requestBody *tview.TextArea) {
	textArea := tview.
		NewTextArea()

	textArea.
		SetBorder(true).
		SetTitle(" {} Editor ").
		SetTitleAlign(tview.AlignLeft)

	return textArea
}

func MethodDropdown(requests []string, handleSelect func(text string, index int), handleInput func(event *tcell.EventKey) *tcell.EventKey) (methodDropdown *tview.DropDown) {
	dropdown := tview.
		NewDropDown().
		SetCurrentOption(0).
		SetOptions(requests, handleSelect).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetTextOptions(" ", "  ", " ", "  ", " "+requests[0]+"  ").
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue),
			tcell.StyleDefault.Background(tcell.ColorBlack),
		)

	dropdown.
		SetBorder(true).
		SetInputCapture(handleInput)

	return dropdown
}

func CreateButton(handleClick func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (createButton *tview.Button) {
	button := tview.
		NewButton("Connect").
		SetLabelColor(tcell.ColorBlue).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetSelectedFunc(handleClick)

	button.
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetInputCapture(handleInput)

	return button
}
