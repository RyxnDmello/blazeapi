package response

import (
	"blazeapi/utils"
	"blazeapi/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeResponse(app *tview.Application) (response *Response, layout *tview.Flex) {
	var clear *tview.Button

	body := widgets.
		NewMessage().
		SetText(utils.INTRODUCTION).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(clear)
				}

				return event
			},
		).
		Render()

	body.
		SetTitle("  Response ").
		SetBorderPadding(0, 0, 0, 0).
		SetTitleAlign(tview.AlignLeft)

	code := widgets.
		NewMessage().
		SetText("Code").
		SetAlignment(tview.AlignCenter).
		Render()

	time := widgets.
		NewMessage().
		SetText("Time").
		SetAlignment(tview.AlignCenter).
		Render()

	status := widgets.
		NewMessage().
		SetText("Status").
		SetAlignment(tview.AlignCenter).
		Render()

	clear = widgets.
		NewButton().
		SetLabel("Clear").
		HandleSelect(
			func() {
				response.Clear()
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(body)
				}

				return event
			},
		).
		Render()

	response = NewResponse().Initialize(body, code, time, status)

	panel := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(response.time, 0, 1, false).
		AddItem(response.status, 0, 1, false).
		AddItem(clear, 0, 1, false)

	layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response.body, 0, 1, true).
		AddItem(panel, 3, 1, true)

	return response, layout
}
