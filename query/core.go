package query

import (
	"strings"

	"blazeapi/core"
	"blazeapi/response"
	"blazeapi/utils"
	"blazeapi/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var METHODS []string = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}

func InitializeQuery(app *tview.Application, response *response.Response) (query *Query, layout *tview.Flex, queryBodyModal *tview.Flex) {
	var create *tview.Button

	method := widgets.
		NewDropdown().
		SetOptions(METHODS).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTab {
					app.SetFocus(query.url)
				}

				return event
			},
		).
		Render()

	url := widgets.
		NewInput().
		SetPlaceholder("Enter Query").
		HandleAcceptance(
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
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(create)
				}

				return event
			},
		).
		Render()

	create = widgets.
		NewButton().
		SetLabel("Create").
		HandleSelect(
			func() {
				request := core.MakeRequest(query.Method(), query.Url(), query.Body())

				response.SetBody(request.Data())
				response.SetTime(request.Time(true))
				response.SetStatus(request.Status())
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(query.method)
				}

				return event
			},
		).
		Render()

	body, queryBodyModal := initializeQueryBodyModal(app)

	query = NewQuery().Initialize(method, url, body)

	layout = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(query.method, 0, 1, false).
		AddItem(query.url, 0, 5, true).
		AddItem(create, 0, 1, false)

	return query, layout, queryBodyModal
}

func initializeQueryBodyModal(app *tview.Application) (body *tview.TextArea, modal *tview.Flex) {
	var format *tview.Button
	var clear *tview.Button

	body = widgets.
		NewTextArea().
		SetPlaceholder("Enter Body").
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(format)
				}

				return event
			},
		).
		Render()

	format = widgets.
		NewButton().
		SetLabel("Format").
		HandleSelect(
			func() {
				body.SetText(utils.Prettier([]byte(body.GetText())), true)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTab {
					app.SetFocus(clear)
				}

				return event
			},
		).
		Render()

	clear = widgets.
		NewButton().
		SetLabel("Clear").
		HandleSelect(
			func() {
				body.SetText("", true)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTab {
					app.SetFocus(body)
				}

				return event
			},
		).
		Render()

	modal = widgets.
		NewModal().
		SetTitle(" Body Data").
		SetDimension(50, 25).
		AddInput(body, true).
		AddButton(format, false).
		AddButton(clear, false).
		Render()

	return body, modal
}
