package query

import (
	"blazeapi/core"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeQuery(app *tview.Application, response *tview.TextView) (*tview.Flex, *tview.Flex) {
	var requestMethod *tview.DropDown
	var requestQuery *tview.InputField
	var requestButton *tview.Button

	var method string
	var body string
	var url string

	requestBody, requestBodyModal := initializeQueryBody()

	requestMethod = MethodDropdown(
		[]string{
			"GET", "POST", "PATCH", "PUT", "DELETE",
		},
		func(text string, index int) {
			method = text
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(requestQuery)
			}

			return event
		},
	)

	requestQuery = RequestInput(
		func(textToCheck string, lastChar rune) bool {
			if !strings.HasPrefix(textToCheck, "http") {
				requestQuery.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			if !strings.Contains(textToCheck, "/") {
				requestQuery.SetFieldTextColor(tcell.ColorRed)
				return true
			}

			url = textToCheck

			requestQuery.SetFieldTextColor(tcell.ColorWhite)

			return true
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(requestButton)
			}

			return event
		},
	)

	requestButton = CreateButton(
		func() {
			body = requestBody.GetText()

			request := core.MakeRequest(method, url, body)

			response.SetText(request.GetTime(true))
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(requestMethod)
			}

			return event
		},
	)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(requestMethod, 0, 1, false).
		AddItem(requestQuery, 0, 5, true).
		AddItem(requestButton, 0, 1, false)

	return flex, requestBodyModal
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
