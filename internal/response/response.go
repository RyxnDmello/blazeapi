package response

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ryxndmello/flame/widgets/button"
	"github.com/ryxndmello/flame/widgets/glyph"
	"github.com/ryxndmello/flame/widgets/message"
)

type Response struct {
	body   *tview.TextView
	time   *tview.TextView
	code   *tview.TextView
	status *tview.TextView
	clear  *tview.Button
	layout *tview.Flex
}

func NewResponse() *Response {
	return &Response{
		body:   nil,
		time:   nil,
		code:   nil,
		status: nil,
		clear:  nil,
		layout: nil,
	}
}

func NewDefaultResponse(app *tview.Application) *Response {
	response := NewResponse().
		UseDefaultBody(app).
		UseDefaultTime().
		UseDefaultCode().
		UseDefaultStatus().
		UseDefaultClear(app).
		UseDefaultLayout()

	return response
}

func (response *Response) SetBody(body *message.Message) *Response {
	response.body = body.
		View()

	return response
}

func (response *Response) UseDefaultBody(app *tview.Application) *Response {
	design := message.
		NewDesign().
		SetRootColor(0x181825).
		SetStyle(0x181825, 0xcdd6f4)

	response.body = message.
		NewMessage().
		SetDesign(design).
		SetAlignment(tview.AlignLeft).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(response.clear)
				}

				return event
			},
		).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 2, y + 1, width, height - 1
		}).
		View()

	return response
}

func (response *Response) SetTime(time *message.Message) *Response {
	response.time = time.
		View()

	return response
}

func (response *Response) UseDefaultTime() *Response {
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

	response.time = message.
		NewMessage().
		SetText("Time").
		SetAlignment(tview.AlignCenter).
		SetDesign(timeDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	return response
}

func (response *Response) SetCode(code *message.Message) *Response {
	response.code = code.
		View()

	return response
}

func (response *Response) UseDefaultCode() *Response {
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

	response.code = message.
		NewMessage().
		SetText("Code").
		SetAlignment(tview.AlignCenter).
		SetDesign(timeDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	return response
}

func (response *Response) SetStatus(status *message.Message) *Response {
	response.status = status.
		View()

	return response
}

func (response *Response) UseDefaultStatus() *Response {
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

	response.status = message.
		NewMessage().
		SetText("Code").
		SetAlignment(tview.AlignCenter).
		SetDesign(timeDesign).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y, width - 1, height
		}).
		View()

	return response
}

func (response *Response) SetClear(clear *button.Button) *Response {
	response.clear = clear.
		View()

	return response
}

func (response *Response) UseDefaultClear(app *tview.Application) *Response {
	clearGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x181825, 0xb4befe).
		SetFocusedStyle(0x181825, 0xb4befe).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	clearDesign := button.
		NewDesign().
		SetRootColor(0xd86d8f).
		SetActiveStyle(0xb65278, 0x1e1e2e).
		SetInactiveStyle(0xd86d8f, 0x1e1e2e).
		AddGlyph(clearGlyph)

	response.clear = button.
		NewButton().
		SetLabel("Clear").
		SetDesign(clearDesign).
		HandleClick(func() {
			response.body.Clear()
		}).
		HandleSwitch(func() {
			app.SetFocus(response.body)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 1, height
		}).
		View()

	return response
}

func (response *Response) SetLayout(layout *tview.Flex) *Response {
	response.layout = layout
	return response
}

func (response *Response) UseDefaultLayout() *Response {
	row := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(response.time, 0, 1, false).
		AddItem(response.code, 0, 1, false).
		AddItem(response.status, 0, 1, false).
		AddItem(response.clear, 0, 1, false)

	response.layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(response.body, 0, 1, true).
		AddItem(row, 1, 1, true)

	return response
}

func (response *Response) GetBody() *tview.TextView {
	return response.body
}

func (response *Response) GetTime() *tview.TextView {
	return response.time
}

func (response *Response) GetCode() *tview.TextView {
	return response.code
}

func (response *Response) GetStatus() *tview.TextView {
	return response.status
}

func (response *Response) GetClear() *tview.Button {
	return response.clear
}

func (response *Response) GetLayout() *tview.Flex {
	return response.layout
}
