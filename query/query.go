package query

import (
	"io"
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeQuery(app *tview.Application, response *tview.TextView) *tview.Flex {
	var requests *tview.DropDown
	var query *tview.InputField
	var create *tview.Button

	requests = RequestDropdown(
		[]string{
			"GET", "POST", "PATCH", "PUT", "DELETE",
		},
		func(text string, index int) {
			if text == "GET" {
				query.SetText("")
			}
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(query)
			}

			return event
		},
	)

	query = QueryInput(
		func(textToCheck string, lastChar rune) bool {
			if !strings.HasPrefix(textToCheck, "http") {
				query.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			if !strings.Contains(textToCheck, "/") {
				query.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			query.SetFieldTextColor(tcell.ColorWhite)

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
			resp, err := http.Get(query.GetText())

			if err != nil {
				response.SetText("Failure")
				return
			}

			body, err := io.ReadAll(resp.Body)

			response.SetText(string(body))
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(requests)
			}

			return event
		},
	)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(requests, 0, 1, true).
		AddItem(query, 0, 5, false).
		AddItem(create, 0, 1, false)

	return flex
}
