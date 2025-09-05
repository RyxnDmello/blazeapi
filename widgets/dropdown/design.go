package dropdown

import (
	"blazeapi/widgets/glyph"

	"github.com/gdamore/tcell/v2"
)

type Design struct {
	glyphs                        []glyph.Glyph
	rootColor                     int32
	fieldBackgroundColor          int32
	fieldForegroundColor          int32
	labelBackgroundColor          int32
	labelForegroundColor          int32
	focusedBackgroundColor        int32
	focusedForegroundColor        int32
	disabledBackgroundColor       int32
	disabledForegroundColor       int32
	listSelectedBackgroundColor   int32
	listSelectedForegroundColor   int32
	listUnselectedBackgroundColor int32
	listUnselectedForegroundColor int32
}

func NewDesign() Design {
	return Design{
		glyphs:                        make([]glyph.Glyph, 0),
		rootColor:                     tcell.ColorDefault.Hex(),
		fieldBackgroundColor:          tcell.ColorDefault.Hex(),
		fieldForegroundColor:          tcell.ColorDefault.Hex(),
		labelBackgroundColor:          tcell.ColorDefault.Hex(),
		labelForegroundColor:          tcell.ColorDefault.Hex(),
		focusedBackgroundColor:        tcell.ColorDefault.Hex(),
		focusedForegroundColor:        tcell.ColorDefault.Hex(),
		disabledBackgroundColor:       tcell.ColorDefault.Hex(),
		disabledForegroundColor:       tcell.ColorDefault.Hex(),
		listSelectedBackgroundColor:   tcell.ColorDefault.Hex(),
		listSelectedForegroundColor:   tcell.ColorDefault.Hex(),
		listUnselectedBackgroundColor: tcell.ColorDefault.Hex(),
		listUnselectedForegroundColor: tcell.ColorDefault.Hex(),
	}
}

func (design Design) SetRootColor(color int32) Design {
	design.rootColor = color
	return design
}

func (design Design) SetFieldStyle(fieldBackgroundColor, fieldForegroundColor int32) Design {
	design.fieldBackgroundColor = fieldBackgroundColor
	design.fieldForegroundColor = fieldForegroundColor
	return design
}

func (design Design) SetLabelStyle(labelBackgroundColor, labelForegroundColor int32) Design {
	design.labelBackgroundColor = labelBackgroundColor
	design.labelForegroundColor = labelForegroundColor
	return design
}

func (design Design) SetFocusedStyle(focusedBackgroundColor, focusedForegroundColor int32) Design {
	design.focusedBackgroundColor = focusedBackgroundColor
	design.focusedForegroundColor = focusedForegroundColor
	return design
}

func (design Design) SetDisabledStyle(disabledBackgroundColor, disabledForegroundColor int32) Design {
	design.disabledBackgroundColor = disabledBackgroundColor
	design.disabledForegroundColor = disabledForegroundColor
	return design
}

func (design Design) SetListSelectedStyle(listSelectedBackgroundColor, listSelectedForegroundColor int32) Design {
	design.listSelectedBackgroundColor = listSelectedBackgroundColor
	design.listSelectedForegroundColor = listSelectedForegroundColor
	return design
}

func (design Design) SetListUnselectedStyle(listUnselectedBackgroundColor, listUnselectedForegroundColor int32) Design {
	design.listUnselectedBackgroundColor = listUnselectedBackgroundColor
	design.listUnselectedForegroundColor = listUnselectedForegroundColor
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

func (design Design) GetFieldStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.fieldBackgroundColor)).
		Foreground(tcell.NewHexColor(design.fieldForegroundColor))

	return style
}

func (design Design) GetLabelStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.labelBackgroundColor)).
		Foreground(tcell.NewHexColor(design.labelForegroundColor))

	return style
}

func (design Design) GetFocusedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.focusedBackgroundColor)).
		Foreground(tcell.NewHexColor(design.focusedForegroundColor))

	return style
}

func (design Design) GetDisabledStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.disabledBackgroundColor)).
		Foreground(tcell.NewHexColor(design.disabledForegroundColor))

	return style
}

func (design Design) GetListSelectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.listSelectedBackgroundColor)).
		Foreground(tcell.NewHexColor(design.listSelectedForegroundColor))

	return style
}

func (design Design) GetListUnselectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.listUnselectedBackgroundColor)).
		Foreground(tcell.NewHexColor(design.listUnselectedForegroundColor))

	return style
}

func (design Design) GetFieldStyleDecompose() (tcell.Color, tcell.Color) {
	fieldBackgroundColor, fieldForegroundColor, _ := design.GetFieldStyle().Decompose()
	return fieldBackgroundColor, fieldForegroundColor
}

func (design Design) GetLabelStyleDecompose() (tcell.Color, tcell.Color) {
	labelBackgroundColor, labelForegroundColor, _ := design.GetLabelStyle().Decompose()
	return labelBackgroundColor, labelForegroundColor
}

func (design Design) GetFocusedStyleDecompose() (tcell.Color, tcell.Color) {
	focusedBackgroundColor, focusedForegroundColor, _ := design.GetFocusedStyle().Decompose()
	return focusedBackgroundColor, focusedForegroundColor
}

func (design Design) GetDisabledStyleDecompose() (tcell.Color, tcell.Color) {
	disabledBackgroundColor, disabledForegroundColor, _ := design.GetDisabledStyle().Decompose()
	return disabledBackgroundColor, disabledForegroundColor
}

func (design Design) GetListSelectedStyleDecompose() (tcell.Color, tcell.Color) {
	listSelectedBackgroundColor, listSelectedForegroundColor, _ := design.GetListSelectedStyle().Decompose()
	return listSelectedBackgroundColor, listSelectedForegroundColor
}

func (design Design) GetListUnelectedStyleDecompose() (tcell.Color, tcell.Color) {
	listUnselectedBackgroundColor, listUnselectedForegroundColor, _ := design.GetListUnselectedStyle().Decompose()
	return listUnselectedBackgroundColor, listUnselectedForegroundColor
}
