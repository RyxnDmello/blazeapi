package query

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RequestInput(handleAccept func(textToCheck string, lastChar rune) bool, handleInput func(event *tcell.EventKey) *tcell.EventKey) *tview.InputField {
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

func RequestBody() *tview.TextArea {
	body := tview.
		NewTextArea()

	body.
		SetBorder(true).
		SetTitle(" {} Editor ").
		SetTitleAlign(tview.AlignLeft)

	return body
}

func MethodDropdown(requests []string, handleSelect func(text string, index int), handleInput func(event *tcell.EventKey) *tcell.EventKey) *tview.DropDown {
	request := tview.
		NewDropDown().
		SetCurrentOption(0).
		SetOptions(requests, handleSelect).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetTextOptions(" ", "  ", " ", "  ", " "+requests[0]+"  ").
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue),
			tcell.StyleDefault.Background(tcell.ColorBlack),
		)

	request.
		SetBorder(true).
		SetInputCapture(handleInput)

	return request
}

func CreateButton(handleSelect func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) *tview.Button {
	create := tview.
		NewButton("Connect").
		SetLabelColor(tcell.ColorBlue).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetSelectedFunc(handleSelect)

	create.
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetInputCapture(handleInput)

	return create
}
