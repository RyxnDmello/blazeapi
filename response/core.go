package response

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeResponse(app *tview.Application) (response Response, layout *tview.Flex) {
	var clear *tview.Button

	response.body = Body(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(clear)
			}

			return event
		},
	)

	response.code = Element("Code")
	response.time = Element("Time")
	response.status = Element("Status")

	clear = Clear(
		func() {
			response.Clear()
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(response.body)
			}

			return event
		},
	)

	analysis := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(response.time, 0, 1, false).
		AddItem(response.status, 0, 1, false).
		AddItem(clear, 0, 1, false)

	layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response.body, 0, 1, true).
		AddItem(analysis, 3, 1, true)

	return response, layout
}
