package response

import (
	"blazeapi/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Body(handleInput func(event *tcell.EventKey) *tcell.EventKey) (body *tview.TextView) {
	body = tview.
		NewTextView().
		SetText(utils.INTRODUCTION)

	body.
		SetBorder(true).
		SetTitle(" 󰅪 Response ").
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(handleInput)

	return body
}

func Element(label string) (element *tview.TextView) {
	element = tview.
		NewTextView().
		SetText(label).
		SetTextAlign(tview.AlignCenter)

	element.
		SetBorder(true)

	return element
}

func Clear(handleClick func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (clear *tview.Button) {
	clear = tview.
		NewButton(" 󰇾 Clear ").
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetSelectedFunc(handleClick)

	clear.
		SetBorder(true).
		SetBackgroundColor(tcell.ColorNone).
		SetInputCapture(handleInput)

	return clear
}
