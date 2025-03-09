package response

import (
	"fmt"

	"blazeapi/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Response struct {
	body   *tview.TextView
	code   *tview.TextView
	time   *tview.TextView
	status *tview.TextView
}

func NewResponse() *Response {
	return &Response{
		body:   nil,
		code:   nil,
		time:   nil,
		status: nil,
	}
}

func (query *Response) Initialize(body *tview.TextView, code *tview.TextView, time *tview.TextView, status *tview.TextView) *Response {
	query.body = body
	query.code = code
	query.time = time
	query.status = status
	return query
}

func (response *Response) SetBody(body string) {
	response.body.SetText(body)
}

func (response *Response) SetCode(code int) {
	color := tcell.ColorWhite

	if code < 300 {
		color = tcell.ColorRed
	}

	response.code.
		SetText(fmt.Sprintf("%d", code)).
		SetTextColor(color).
		SetBorderColor(color)
}

func (response *Response) SetTime(time string) {
	response.time.SetText(fmt.Sprintf("%sms", time))
}

func (response *Response) SetStatus(status string) {
	response.status.SetText(status)
}

func (response *Response) Clear() {
	response.body.
		SetText(utils.INTRODUCTION).
		SetTextColor(tcell.ColorWhite).
		SetBorderColor(tcell.ColorWhite)

	response.code.
		SetText("Code").
		SetTextColor(tcell.ColorWhite).
		SetBorderColor(tcell.ColorWhite)

	response.time.
		SetText("Time").
		SetTextColor(tcell.ColorWhite).
		SetBorderColor(tcell.ColorWhite)

	response.status.
		SetText("Status").
		SetTextColor(tcell.ColorWhite).
		SetBorderColor(tcell.ColorWhite)
}
