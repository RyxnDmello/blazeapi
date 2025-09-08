package area

import (
	"github.com/ryxndmello/flame/widgets/glyph"

	"github.com/gdamore/tcell/v2"
)

type Design struct {
	glyphs                     []glyph.Glyph
	rootColor                  int32
	textBackgroundColor        int32
	textForegroundColor        int32
	labelBackgroundColor       int32
	labelForegroundColor       int32
	placeholderBackgroundColor int32
	placeholderForegroundColor int32
}

func NewDesign() Design {
	return Design{
		glyphs:                     make([]glyph.Glyph, 0),
		rootColor:                  tcell.ColorDefault.Hex(),
		textBackgroundColor:        tcell.ColorDefault.Hex(),
		textForegroundColor:        tcell.ColorDefault.Hex(),
		labelBackgroundColor:       tcell.ColorDefault.Hex(),
		labelForegroundColor:       tcell.ColorDefault.Hex(),
		placeholderBackgroundColor: tcell.ColorDefault.Hex(),
		placeholderForegroundColor: tcell.ColorDefault.Hex(),
	}
}

func (design Design) SetRootColor(color int32) Design {
	design.rootColor = color
	return design
}

func (design Design) SetTextStyle(textBackgroundColor, textForegroundColor int32) Design {
	design.textBackgroundColor = textBackgroundColor
	design.textForegroundColor = textForegroundColor
	return design
}

func (design Design) SetLabelStyle(labelBackgroundColor, labelForegroundColor int32) Design {
	design.labelBackgroundColor = labelBackgroundColor
	design.labelForegroundColor = labelForegroundColor
	return design
}

func (design Design) SetPlaceholderStyle(placeholderBackgroundColor, placeholderForegroundColor int32) Design {
	design.placeholderBackgroundColor = placeholderBackgroundColor
	design.placeholderForegroundColor = placeholderForegroundColor
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

func (design Design) GetTextStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.textBackgroundColor)).
		Foreground(tcell.NewHexColor(design.textForegroundColor))

	return style
}

func (design Design) GetLabelStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.labelBackgroundColor)).
		Foreground(tcell.NewHexColor(design.labelForegroundColor))

	return style
}

func (design Design) GetPlaceholderStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.placeholderBackgroundColor)).
		Foreground(tcell.NewHexColor(design.placeholderForegroundColor))

	return style
}

func (design Design) GetTextStyleDecompose() (tcell.Color, tcell.Color) {
	textBackgroundColor, textForegroundColor, _ := design.GetTextStyle().Decompose()
	return textBackgroundColor, textForegroundColor
}

func (design Design) GetLabelStyleDecompose() (tcell.Color, tcell.Color) {
	labelBackgroundColor, labelForegroundColor, _ := design.GetLabelStyle().Decompose()
	return labelBackgroundColor, labelForegroundColor
}

func (design Design) GetPlaceholderStyleDecompose() (tcell.Color, tcell.Color) {
	placeholderBackgroundColor, placeholderForegroundColor, _ := design.GetPlaceholderStyle().Decompose()
	return placeholderBackgroundColor, placeholderForegroundColor
}
