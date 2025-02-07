package response

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Response struct {
	body   *tview.TextView
	code   *tview.TextView
	time   *tview.TextView
	status *tview.TextView
}

func (response *Response) SetBody(body string) {
	response.body.SetText(body)
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
	response.time.SetText(fmt.Sprintf("%sms", time))
}

func (response *Response) SetStatus(status string) {
	response.status.SetText(status)
}

func (response *Response) Clear() {
	response.body.Clear()
	response.time.Clear().SetText("Time")
	response.status.Clear().SetText("Status")
}
