package glyph

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Glyph struct {
	symbol                  rune
	backgroundColor         int32
	foregroundColor         int32
	focusedBackgroundColor  int32
	focusedForegroundColor  int32
	disabledBackgroundColor int32
	disabledForegroundColor int32
	handlePosition          func(elementX, elementY, elementWidth, elementHeight int) (int, int)
	handleFocusedPosition   func(elementX, elementY, elementWidth, elementHeight int) (int, int)
	handleDisabledPosition  func(elementX, elementY, elementWidth, elementHeight int) (int, int)
}

func NewGlyph() Glyph {
	return Glyph{
		symbol:                  ' ',
		backgroundColor:         tcell.ColorDefault.Hex(),
		foregroundColor:         tcell.ColorDefault.Hex(),
		focusedBackgroundColor:  tcell.ColorDefault.Hex(),
		focusedForegroundColor:  tcell.ColorDefault.Hex(),
		disabledBackgroundColor: tcell.ColorDefault.Hex(),
		disabledForegroundColor: tcell.ColorDefault.Hex(),
		handlePosition:          nil,
		handleFocusedPosition:   nil,
		handleDisabledPosition:  nil,
	}
}

func (glyph Glyph) SetSymbol(symbol rune) Glyph {
	glyph.symbol = symbol
	return glyph
}

func (glyph Glyph) SetStyle(backgroundColor, foregroundColor int32) Glyph {
	glyph.backgroundColor = backgroundColor
	glyph.foregroundColor = foregroundColor
	return glyph
}

func (glyph Glyph) SetFocusedStyle(focusedBackgroundColor, focusedForegroundColor int32) Glyph {
	glyph.focusedBackgroundColor = focusedBackgroundColor
	glyph.focusedForegroundColor = focusedForegroundColor
	return glyph
}

func (glyph Glyph) SetDisabledStyle(disabledBackgroundColor, disabledForegroundColor int32) Glyph {
	glyph.disabledBackgroundColor = disabledBackgroundColor
	glyph.disabledForegroundColor = disabledForegroundColor
	return glyph
}

func (glyph Glyph) HandlePosition(handlePosition func(elementX, elementY, elementWidth, elementHeight int) (int, int)) Glyph {
	glyph.handlePosition = handlePosition
	glyph.handleFocusedPosition = handlePosition
	glyph.handleDisabledPosition = handlePosition
	return glyph
}

func (glyph Glyph) HandleFocusedPosition(handleFocusedPosition func(elementX, elementY, elementWidth, elementHeight int) (int, int)) Glyph {
	glyph.handleFocusedPosition = handleFocusedPosition
	return glyph
}

func (glyph Glyph) HandleDisabledPosition(handleDisabledPosition func(elementX, elementY, elementWidth, elementHeight int) (int, int)) Glyph {
	glyph.handleDisabledPosition = handleDisabledPosition
	return glyph
}

func (glyph Glyph) GetSymbol() rune {
	return glyph.symbol
}

func (glyph Glyph) GetStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(glyph.backgroundColor)).
		Foreground(tcell.NewHexColor(glyph.foregroundColor))

	return style
}

func (glyph Glyph) GetFocusedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(glyph.focusedBackgroundColor)).
		Foreground(tcell.NewHexColor(glyph.focusedForegroundColor))

	return style
}

func (glyph Glyph) GetDisabledStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(glyph.disabledBackgroundColor)).
		Foreground(tcell.NewHexColor(glyph.disabledForegroundColor))

	return style
}

func (glyph Glyph) GetPosition(elementX, elementY, elementWidth, elementHeight int) (int, int) {
	return glyph.handlePosition(elementX, elementY, elementWidth, elementHeight)
}

func (glyph Glyph) GetFocusedPosition(elementX, elementY, elementWidth, elementHeight int) (int, int) {
	return glyph.handleFocusedPosition(elementX, elementY, elementWidth, elementHeight)
}

func (glyph Glyph) GetDisabledPosition(elementX, elementY, elementWidth, elementHeight int) (int, int) {
	return glyph.handleDisabledPosition(elementX, elementY, elementWidth, elementHeight)
}

func (glyph Glyph) Render(element *tview.Box, screen tcell.Screen, focused, disabled bool) {
	if glyph.handleDisabledPosition != nil && disabled {
		x, y := glyph.handleDisabledPosition(element.GetRect())
		screen.SetContent(x, y, glyph.symbol, nil, glyph.GetDisabledStyle())
		return
	}

	if glyph.handleFocusedPosition != nil && focused {
		x, y := glyph.handleFocusedPosition(element.GetRect())
		screen.SetContent(x, y, glyph.symbol, nil, glyph.GetFocusedStyle())
		return
	}

	if glyph.handlePosition != nil {
		x, y := glyph.handlePosition(element.GetRect())
		screen.SetContent(x, y, glyph.symbol, nil, glyph.GetStyle())
	}
}

func (glyph Glyph) RenderGlobally(screen tcell.Screen, x, y int, focused bool, disabled bool) {
	if disabled {
		screen.SetContent(x, y, glyph.symbol, nil, glyph.GetDisabledStyle())
		return
	}

	if focused {
		screen.SetContent(x, y, glyph.symbol, nil, glyph.GetFocusedStyle())
		return
	}

	screen.SetContent(x, y, glyph.symbol, nil, glyph.GetStyle())
}
