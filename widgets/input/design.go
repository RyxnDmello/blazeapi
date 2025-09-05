package input

import (
	"blazeapi/widgets/glyph"

	"github.com/gdamore/tcell/v2"
)

type Design struct {
	glyphs                                []glyph.Glyph
	rootColor                             int32
	fieldBackgroundColor                  int32
	fieldForegroundColor                  int32
	labelBackgroundColor                  int32
	labelForegroundColor                  int32
	placeholderBackgroundColor            int32
	placeholderForegroundColor            int32
	autocompleteSelectedBackgroundColor   int32
	autocompleteSelectedForegroundColor   int32
	autocompleteUnselectedBackgroundColor int32
	autocompleteUnselectedForegroundColor int32
}

func NewDesign() Design {
	return Design{
		glyphs:                                make([]glyph.Glyph, 0),
		rootColor:                             tcell.ColorDefault.Hex(),
		fieldBackgroundColor:                  tcell.ColorDefault.Hex(),
		fieldForegroundColor:                  tcell.ColorDefault.Hex(),
		labelBackgroundColor:                  tcell.ColorDefault.Hex(),
		labelForegroundColor:                  tcell.ColorDefault.Hex(),
		placeholderBackgroundColor:            tcell.ColorDefault.Hex(),
		placeholderForegroundColor:            tcell.ColorDefault.Hex(),
		autocompleteSelectedBackgroundColor:   tcell.ColorDefault.Hex(),
		autocompleteSelectedForegroundColor:   tcell.ColorDefault.Hex(),
		autocompleteUnselectedBackgroundColor: tcell.ColorDefault.Hex(),
		autocompleteUnselectedForegroundColor: tcell.ColorDefault.Hex(),
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

func (design Design) SetPlaceholderStyle(placeholderBackgroundColor, placeholderForegroundColor int32) Design {
	design.placeholderBackgroundColor = placeholderBackgroundColor
	design.placeholderForegroundColor = placeholderForegroundColor
	return design
}

func (design Design) SetAutocompleteSelectedStyle(autocompleteSelectedBackgroundColor, autocompleteSelectedForegroundColor int32) Design {
	design.autocompleteSelectedBackgroundColor = autocompleteSelectedBackgroundColor
	design.autocompleteSelectedForegroundColor = autocompleteSelectedForegroundColor
	return design
}

func (design Design) SetAutocompleteUnselectedStyle(autocompleteUnselectedBackgroundColor, autocompleteUnselectedForegroundColor int32) Design {
	design.autocompleteUnselectedBackgroundColor = autocompleteUnselectedBackgroundColor
	design.autocompleteUnselectedForegroundColor = autocompleteUnselectedForegroundColor
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

func (design Design) GetPlaceholderStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.placeholderBackgroundColor)).
		Foreground(tcell.NewHexColor(design.placeholderForegroundColor))

	return style
}

func (design Design) GetAutocompleteSelectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.autocompleteSelectedBackgroundColor)).
		Foreground(tcell.NewHexColor(design.autocompleteSelectedForegroundColor))

	return style
}

func (design Design) GetAutocompleteUnselectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.autocompleteUnselectedBackgroundColor)).
		Foreground(tcell.NewHexColor(design.autocompleteUnselectedForegroundColor))

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

func (design Design) GetPlaceholderStyleDecompose() (tcell.Color, tcell.Color) {
	placeholderBackgroundColor, placeholderForegroundColor, _ := design.GetPlaceholderStyle().Decompose()
	return placeholderBackgroundColor, placeholderForegroundColor
}

func (design Design) GetAutocompleteSelectedStyleDecompose() (tcell.Color, tcell.Color) {
	autocompleteSelectedBackgroundColor, autocompleteSelectedForegroundColor, _ := design.GetAutocompleteSelectedStyle().Decompose()
	return autocompleteSelectedBackgroundColor, autocompleteSelectedForegroundColor
}

func (design Design) GetAutocompleteUnselectedStyleDecompose() (tcell.Color, tcell.Color) {
	autocompleteUnselectedBackgroundColor, autocompleteUnselectedForegroundColor, _ := design.GetAutocompleteUnselectedStyle().Decompose()
	return autocompleteUnselectedBackgroundColor, autocompleteUnselectedForegroundColor
}
