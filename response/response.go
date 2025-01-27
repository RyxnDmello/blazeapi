package response

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeResponse() (*tview.TextView, *tview.Flex) {
	response := tview.
		NewTextView()

	response.
		SetBorder(true).
		SetTitle(" 󰅪 Response ").
		SetTitleAlign(tview.AlignLeft)

	clearButton := tview.
		NewButton("Clear").
		SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
		SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))

	clearButton.
		SetBorder(true).
		SetBackgroundColor(tcell.ColorNone)

	tabsLayout := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(nil, 0, 1, false).
		AddItem(clearButton, 0, 1, false)

	responseLayout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response, 0, 1, true).
		AddItem(tabsLayout, 3, 1, true)

	return response, responseLayout
}
