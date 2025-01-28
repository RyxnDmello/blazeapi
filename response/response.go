package response

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeResponse(app *tview.Application) (*tview.TextView, *tview.Flex) {
	var responseBody *tview.TextView
	var clearButton *tview.Button

	responseBody = ResponseBody(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(clearButton)
			}

			return event
		},
	)

	clearButton = ClearButton(
		func() {
			responseBody.Clear()
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(responseBody)
			}

			return event
		},
	)

	tabsLayout := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(nil, 0, 1, false).
		AddItem(clearButton, 0, 1, false)

	responseLayout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(responseBody, 0, 1, true).
		AddItem(tabsLayout, 3, 1, true)

	return responseBody, responseLayout
}
