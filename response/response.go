package response

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Response struct {
	body   *tview.TextView
	status *tview.TextView
	code   *tview.TextView
	time   *tview.TextView
}

func (response *Response) SetBody(body string) {
	response.body.SetText(body).SetToggleHighlights(true)
}

func (response *Response) SetStatus(status string) {
	response.status.SetText(status)
}

func (response *Response) SetCode(code int) {
	response.code.SetText(fmt.Sprintf("%d", code))

	if code < 300 {
		response.code.SetBorderColor(tcell.ColorWhite)
		return
	}

	response.code.SetTextColor(tcell.ColorRed)
	response.code.SetBorderColor(tcell.ColorRed)
}

func (response *Response) SetTime(time string) {
	response.time.SetText(time + "ms")
}

func (response *Response) Clear() {
	response.body.Clear()
	response.time.Clear().SetText("Time")
	response.status.Clear().SetText("Status")
}

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

	details := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(response.time, 0, 1, false).
		AddItem(response.status, 0, 1, false).
		AddItem(clear, 0, 1, false)

	layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response.body, 0, 1, true).
		AddItem(details, 3, 1, true)

	return response, layout
}
