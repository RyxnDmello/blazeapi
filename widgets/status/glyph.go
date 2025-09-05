package status

import (
	"github.com/gdamore/tcell/v2"
)

type Glyph struct {
	symbol          rune
	backgroundColor int32
	foregroundColor int32
}

func NewGlyph() Glyph {
	return Glyph{
		symbol:          rune(0),
		backgroundColor: tcell.ColorDefault.Hex(),
		foregroundColor: tcell.ColorDefault.Hex(),
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

func (glyph Glyph) GetStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(glyph.backgroundColor)).
		Foreground(tcell.NewHexColor(glyph.foregroundColor))

	return style
}
