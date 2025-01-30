package query

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MethodDropdown(requests []string, handleInput func(event *tcell.EventKey) *tcell.EventKey) (method *tview.DropDown) {
	method = tview.
		NewDropDown().
		SetCurrentOption(0).
		SetOptions(requests, nil).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetTextOptions(" ", "  ", " ", "  ", " "+requests[0]+"  ").
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue),
			tcell.StyleDefault.Background(tcell.ColorBlack),
		)

	method.
		SetBorder(true).
		SetInputCapture(handleInput)

	return method
}

func RequestUrl(handleAccept func(textToCheck string, lastChar rune) bool, handleInput func(event *tcell.EventKey) *tcell.EventKey) (url *tview.InputField) {
	url = tview.
		NewInputField().
		SetPlaceholder("Enter URL").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetAcceptanceFunc(handleAccept)

	url.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetBackgroundColor(tcell.ColorBlack).
		SetInputCapture(handleInput)

	return url
}

func RequestBody() (body *tview.TextArea) {
	body = tview.
		NewTextArea()

	body.
		SetBorder(true).
		SetTitle(" {} Editor ").
		SetTitleAlign(tview.AlignLeft)

	return body
}

func CreateButton(handleClick func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (create *tview.Button) {
	create = tview.
		NewButton("Connect").
		SetLabelColor(tcell.ColorBlue).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true)).
		SetSelectedFunc(handleClick)

	create.
		SetBorder(true).
		SetBorderColor(tcell.ColorBlue).
		SetInputCapture(handleInput)

	return create
}
