package query

import (
	"strings"

	"blazeapi/core"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeQuery(app *tview.Application, response response.Response) (query Query, layout *tview.Flex, bodyModal *tview.Flex) {
	var create *tview.Button

	query.method = Method(
		methods,
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(query.url)
			}

			return event
		},
	)

	query.url = Url(
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

	query.body, bodyModal = queryBody()

	create = Create(
		func() {
			request := core.MakeRequest(query.Method(), query.Url(), query.Body())

			response.SetBody(request.Data())
			response.SetTime(request.Time(true))
			response.SetStatus(request.Status())
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

	return query, layout, bodyModal
}

func queryBody() (body *tview.TextArea, bodyModal *tview.Flex) {
	body = Body()

	alignment := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(body, 15, 1, true).
		AddItem(nil, 0, 1, false)

	bodyModal = tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 60, 1, true).
		AddItem(nil, 0, 1, false)

	return body, bodyModal
}
