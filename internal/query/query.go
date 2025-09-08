package query

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ryxndmello/flame/internal/response"

	"github.com/ryxndmello/flame/widgets/button"
	"github.com/ryxndmello/flame/widgets/dropdown"
	"github.com/ryxndmello/flame/widgets/glyph"
	"github.com/ryxndmello/flame/widgets/input"
)

var methods = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}

type Query struct {
	url    *tview.InputField
	method *tview.DropDown
	create *tview.Button
	layout *tview.Flex
}

func NewQuery() *Query {
	return &Query{
		url:    nil,
		method: nil,
		create: nil,
		layout: nil,
	}
}

func NewDefaultQuery(app *tview.Application, response *response.Response) *Query {
	query := NewQuery().
		UseDefaultMethod(app).
		UseDefaultUrl(app, response).
		UseDefaultCreate(app).
		UseDefaultLayout()

	return query
}

func (query *Query) SetUrl(url *input.Input) *Query {
	query.url = url.
		View()

	return query
}

func (query *Query) UseDefaultUrl(app *tview.Application, response *response.Response) *Query {
	design := input.
		NewDesign().
		SetRootColor(0x1e1e2e).
		SetFieldStyle(0x1e1e2e, 0xcdd6f4).
		SetLabelStyle(0x1e1e2e, 0xcdd6f4).
		SetPlaceholderStyle(0x1e1e2e, 0xcdd6f4).
		SetAutocompleteSelectedStyle(0x3a6cb5, 0xf8f8f2).
		SetAutocompleteUnselectedStyle(0x2c558f, 0xf8f8f2)

	query.url = input.
		NewInput().
		SetDesign(design).
		SetPlaceholder("Enter Request").
		HandleClick(func() {
		}).
		HandleSwitch(func() {
			app.SetFocus(query.create)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	return query
}

func (query *Query) SetMethod(method *dropdown.Dropdown) *Query {
	query.method = method.
		View()

	return query
}

func (query *Query) UseDefaultMethod(app *tview.Application) *Query {
	glyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x1e1e2e, 0x89b4fa).
		SetFocusedStyle(0x1e1e2e, 0x89b4fa).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	design := dropdown.
		NewDesign().
		SetRootColor(0x89b4fa).
		SetLabelStyle(0x89b4fa, 0x1e1e2e).
		SetFieldStyle(0x89b4fa, 0x1e1e2e).
		SetFocusedStyle(0x89b4fa, 0x1e1e2e).
		SetListSelectedStyle(0x3a6cb5, 0xf8f8f2).
		SetListUnselectedStyle(0x2c558f, 0xf8f8f2).
		AddGlyph(glyph)

	query.method = dropdown.
		NewDropdown().
		SetDesign(design).
		SetOptions(methods).
		SetPrefix(" ").
		SetSuffix("").
		SetListPrefix(" ").
		SetListSuffix("       ").
		HandleSwitch(func() {
			app.SetFocus(query.url)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	return query
}

func (query *Query) SetCreate(create *button.Button) *Query {
	query.create = create.
		View()

	return query
}

func (query *Query) UseDefaultCreate(app *tview.Application) *Query {
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

	query.create = button.
		NewButton().
		SetLabel("Create").
		SetDesign(createDesign).
		HandleClick(func() {
		}).
		HandleSwitch(func() {
			app.SetFocus(query.method)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	return query
}

func (query *Query) SetLayout(layout *tview.Flex) *Query {
	query.layout = layout
	return query
}

func (query *Query) UseDefaultLayout() *Query {
	query.layout = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(query.method, 15, 1, false).
		AddItem(query.url, 0, 1, true).
		AddItem(query.create, 20, 1, false)

	return query
}

func (query *Query) GetUrl() *tview.InputField {
	return query.url
}

func (query *Query) GetMethod() *tview.DropDown {
	return query.method
}

func (query *Query) GetCreate() *tview.Button {
	return query.create
}

func (query *Query) GetLayout() *tview.Flex {
	return query.layout
}
