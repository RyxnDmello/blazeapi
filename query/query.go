package query

import (
	"slices"
	"strings"

	"blazeapi/core"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var options []string = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}

type Query struct {
	method *tview.DropDown
	url    *tview.InputField
	body   *tview.TextArea
}

func (query *Query) GetMethod() string {
	_, method := query.method.GetCurrentOption()
	return method
}

func (query *Query) SetMethod(method string) {
	if index := slices.Index(options, method); index != -1 {
		query.method.SetCurrentOption(index)
		return
	}

	query.method.SetCurrentOption(0)
}

func (query *Query) GetUrl() string {
	return query.url.GetText()
}

func (query *Query) SetUrl(url string) {
	query.url.SetText(url)
}

func (query *Query) GetBody() string {
	return query.body.GetText()
}

func (query *Query) SetBody(body string) {
	query.body.SetText(body, false)
}

func InitializeQuery(app *tview.Application, response response.Response) (query Query, layout *tview.Flex, modal *tview.Flex) {
	var create *tview.Button

	query.body, modal = initializeQueryBody()

	query.method = MethodDropdown(
		options,
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
