package response

import (
	"blazeapi/widgets/button"
	"blazeapi/widgets/glyph"
	"blazeapi/widgets/message"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeResponse(app *tview.Application) (response *Response, layout *tview.Flex) {
	var clear *tview.Button

	bodyDesign := message.
		NewDesign().
		SetRootColor(0x181825).
		SetStyle(0x181825, 0xcdd6f4)

	body := message.
		NewMessage().
		SetAlignment(tview.AlignLeft).
		SetDesign(bodyDesign).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(clear)
				}

				return event
			},
		).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 2, y + 1, width, height - 1
		}).
		View()

	timeGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	timeDesign := message.
		NewDesign().
		SetRootColor(0xb4befe).
		SetStyle(0xb4befe, 0x313244).
		AddGlyph(timeGlyph)

	time := message.
		NewMessage().
		SetText("Time").
		SetAlignment(tview.AlignCenter).
		SetDesign(timeDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	codeStartGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	codeEndGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	codeDesign := message.
		NewDesign().
		SetRootColor(0xb4befe).
		SetStyle(0xb4befe, 0x313244).
		AddGlyph(codeStartGlyph).
		AddGlyph(codeEndGlyph)

	code := message.
		NewMessage().
		SetText("Code").
		SetAlignment(tview.AlignCenter).
		SetDesign(codeDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 2, height
		}).
		View()

	statusStartGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	statusEndGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	statusDesign := message.
		NewDesign().
		SetRootColor(0xb4befe).
		SetStyle(0xb4befe, 0x313244).
		AddGlyph(statusStartGlyph).
		AddGlyph(statusEndGlyph)

	status := message.
		NewMessage().
		SetText("Status").
		SetAlignment(tview.AlignCenter).
		SetDesign(statusDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 2, height
		}).
		View()

	clearGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x181825, 0xd86d8f).
		SetFocusedStyle(0x181825, 0xb65278).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	clearDesign := button.
		NewDesign().
		SetRootColor(0xd86d8f).
		SetActiveStyle(0xb65278, 0x1e1e2e).
		SetInactiveStyle(0xd86d8f, 0x1e1e2e).
		AddGlyph(clearGlyph)

	clear = button.
		NewButton().
		SetLabel("Clear").
		SetDesign(clearDesign).
		HandleClick(func() {
			body.Clear()
		}).
		HandleSwitch(func() {
			app.SetFocus(body)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	response = NewResponse().Initialize(body, code, time, status)

	panel := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(response.time, 0, 1, false).
		AddItem(response.code, 0, 1, false).
		AddItem(response.status, 0, 1, false).
		AddItem(clear, 0, 1, false)

	layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response.body, 0, 1, true).
		AddItem(panel, 1, 1, true)

	return response, layout
}
