package response

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ResponseBody(handleInput func(event *tcell.EventKey) *tcell.EventKey) (responseBody *tview.TextView) {
	textView := tview.
		NewTextView()

	textView.
		SetBorder(true).
		SetTitle(" 󰅪 Response ").
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(handleInput)

	return textView
}

func ClearButton(handleClick func(), handleInput func(event *tcell.EventKey) *tcell.EventKey) (clearButton *tview.Button) {
	clearButton = tview.
		NewButton("Clear").
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetSelectedFunc(handleClick)

	clearButton.
		SetBorder(true).
		SetBackgroundColor(tcell.ColorNone).
		SetInputCapture(handleInput)

	return clearButton
}
