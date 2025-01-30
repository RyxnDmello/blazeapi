package query

import (
	"strings"

	"blazeapi/core"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Query struct {
	method *tview.DropDown
	url    *tview.InputField
	body   *tview.TextArea
}

func (query *Query) GetMethod() string {
	_, method := query.method.GetCurrentOption()
	return method
}

func (query *Query) GetUrl() string {
	return query.url.GetText()
}

func (query *Query) GetBody() string {
	return query.body.GetText()
}

func InitializeQuery(app *tview.Application, response response.Response) (query Query, layout *tview.Flex, modal *tview.Flex) {
	var create *tview.Button

	query.body, modal = initializeQueryBody()

	query.method = MethodDropdown(
		[]string{
			"GET", "POST", "PATCH", "PUT", "DELETE",
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(query.url)
			}

			return event
		},
	)

	query.url = RequestUrl(
		func(textToCheck string, lastChar rune) bool {
			if !strings.HasPrefix(textToCheck, "http") {
				query.url.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			if !strings.Contains(textToCheck, "/") {
				query.url.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			query.url.SetFieldTextColor(tcell.ColorWhite)

			return true
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(create)
			}

			return event
		},
	)

	create = CreateButton(
		func() {
			request := core.MakeRequest(query.GetMethod(), query.GetUrl(), query.GetBody())

			response.SetBody(request.GetData())
			response.SetTime(request.GetTime(true))
			response.SetCode(request.GetCode())
			response.SetStatus(request.GetStatus())
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(query.method)
			}

			return event
		},
	)

	layout = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(query.method, 0, 1, false).
		AddItem(query.url, 0, 5, true).
		AddItem(create, 0, 1, false)

	return query, layout, modal
}

func initializeQueryBody() (*tview.TextArea, *tview.Flex) {
	requestBody := RequestBody()

	alignment := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(requestBody, 15, 1, true).
		AddItem(nil, 0, 1, false)

	requestBodyModal := tview.
		NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 60, 1, true).
		AddItem(nil, 0, 1, false)

	return requestBody, requestBodyModal
}
