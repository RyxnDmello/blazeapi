package message

import (
	"github.com/ryxndmello/flame/widgets/glyph"

	"github.com/gdamore/tcell/v2"
)

type Design struct {
	glyphs          []glyph.Glyph
	rootColor       int32
	backgroundColor int32
	foregroundColor int32
}

func NewDesign() Design {
	return Design{
		glyphs:          make([]glyph.Glyph, 0),
		rootColor:       tcell.ColorDefault.Hex(),
		backgroundColor: tcell.ColorDefault.Hex(),
		foregroundColor: tcell.ColorDefault.Hex(),
	}
}

func (design Design) SetRootColor(color int32) Design {
	design.rootColor = color
	return design
}

func (design Design) SetStyle(backgroundColor, foregroundColor int32) Design {
	design.backgroundColor = backgroundColor
	design.foregroundColor = foregroundColor
	return design
}

func (design Design) AddGlyph(glyph glyph.Glyph) Design {
	design.glyphs = append(design.glyphs, glyph)
	return design
}

func (design Design) GetGlyph(index int) glyph.Glyph {
	length := len(design.glyphs)

	if index+1 > length {
		return design.glyphs[length-1]
	}

	return design.glyphs[index]
}

func (design Design) GetGlyphs() []glyph.Glyph {
	return design.glyphs
}

func (design Design) GetRootColor() tcell.Color {
	rootColor := tcell.NewHexColor(design.rootColor)
	return rootColor
}

func (design Design) GetStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.backgroundColor)).
		Foreground(tcell.NewHexColor(design.foregroundColor))

	return style
}

func (design Design) GetStyleDecompose() (tcell.Color, tcell.Color) {
	backgroundColor, foregroundColor, _ := design.GetStyle().Decompose()
	return backgroundColor, foregroundColor
}
