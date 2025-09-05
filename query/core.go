package query

import (
	"blazeapi/core"
	"blazeapi/response"
	"blazeapi/utils"

	"blazeapi/widgets/button"
	"blazeapi/widgets/dropdown"
	"blazeapi/widgets/glyph"
	"blazeapi/widgets/input"

	ember "github.com/ryxndmello/blazelib/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var METHODS []string = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}

func InitializeQuery(app *tview.Application, response *response.Response) (query *Query, layout *tview.Flex, queryBodyModal *tview.Flex) {
	var create *tview.Button
	var url *tview.InputField
	var method *tview.DropDown

	methodGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x1e1e2e, 0x89b4fa).
		SetFocusedStyle(0x1e1e2e, 0x89b4fa).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	methodDesign := dropdown.
		NewDesign().
		SetRootColor(0x89b4fa).
		SetLabelStyle(0x89b4fa, 0x1e1e2e).
		SetFieldStyle(0x89b4fa, 0x1e1e2e).
		SetFocusedStyle(0x89b4fa, 0x1e1e2e).
		SetListSelectedStyle(0x3a6cb5, 0xf8f8f2).
		SetListUnselectedStyle(0x2c558f, 0xf8f8f2).
		AddGlyph(methodGlyph)

	method = dropdown.
		NewDropdown().
		SetOptions(METHODS).
		SetPrefix(" ").
		SetSuffix("").
		SetListPrefix(" ").
		SetListSuffix("       ").
		SetDesign(methodDesign).
		HandleSwitch(func() {
			app.SetFocus(url)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	urlDesign := input.
		NewDesign().
		SetRootColor(0x1e1e2e).
		SetFieldStyle(0x1e1e2e, 0xcdd6f4).
		SetLabelStyle(0x1e1e2e, 0xcdd6f4).
		SetPlaceholderStyle(0x1e1e2e, 0xcdd6f4).
		SetAutocompleteSelectedStyle(0x3a6cb5, 0xf8f8f2).
		SetAutocompleteUnselectedStyle(0x2c558f, 0xf8f8f2)

	urlAutocomplete := input.
		NewAutocomplete().
		SetPadding(1, 5).
		AddEntry("{{host}}", []string{"http://localhost:{{port}}", "https://{{domain}}"})

	url = input.
		NewInput().
		SetDesign(urlDesign).
		SetPlaceholder("Enter Request").
		SetAutocomplete(urlAutocomplete).
		HandleClick(func() {
			response.SetBody(url.GetText())
		}).
		HandleSwitch(func() {
			app.SetFocus(create)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	createGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x1e1e2e, 0xa6e3a1).
		SetFocusedStyle(0x1e1e2e, 0x8acb86).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	createDesign := button.
		NewDesign().
		SetRootColor(0xa6e3a1).
		SetActiveStyle(0x8acb86, 0x1e1e2e).
		SetInactiveStyle(0xa6e3a1, 0x1e1e2e).
		AddGlyph(createGlyph)

	create = button.
		NewButton().
		SetLabel("Create").
		SetDesign(createDesign).
		HandleClick(func() {
			request := core.NewRequest().MakeRequest(query.Method(), query.Url(), query.Body())
			response.SetBody(request.Data())
			response.SetCode(request.Code())
			response.SetTime(request.Time(true))
			response.SetStatus(request.Status())
		}).
		HandleSwitch(func() {
			app.SetFocus(method)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	body, queryBodyModal := outlineModal(app)

	query = NewQuery().Initialize(method, url, body)

	layout = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(query.method, 15, 1, false).
		AddItem(query.url, 0, 1, true).
		AddItem(create, 20, 1, false)

	return query, layout, queryBodyModal
}

func outlineModal(app *tview.Application) (body *tview.TextArea, modal *tview.Flex) {
	var format *tview.Button
	var clear *tview.Button

	body = ember.
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

	format = ember.
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

	clear = ember.
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

	modal = ember.
		NewModal().
		SetTitle(" Body Data").
		SetDimension(50, 25).
		AddInput(body, true).
		AddButton(format, false).
		AddButton(clear, false).
		Render()

	return body, modal
}
