package button

import (
	"blazeapi/widgets/glyph"

	"github.com/gdamore/tcell/v2"
)

type Design struct {
	glyphs                  []glyph.Glyph
	rootColor               int32
	activeBackgroundColor   int32
	activeForegroundColor   int32
	inactiveBackgroundColor int32
	inactiveForegroundColor int32
	disabledBackgroundColor int32
	disabledForegroundColor int32
}

func NewDesign() Design {
	return Design{
		glyphs:                  make([]glyph.Glyph, 0),
		rootColor:               tcell.ColorDefault.Hex(),
		activeBackgroundColor:   tcell.ColorDefault.Hex(),
		activeForegroundColor:   tcell.ColorDefault.Hex(),
		inactiveBackgroundColor: tcell.ColorDefault.Hex(),
		inactiveForegroundColor: tcell.ColorDefault.Hex(),
		disabledBackgroundColor: tcell.ColorDefault.Hex(),
		disabledForegroundColor: tcell.ColorDefault.Hex(),
	}
}

func (design Design) SetRootColor(color int32) Design {
	design.rootColor = color
	return design
}

func (design Design) SetActiveStyle(activeBackgroundColor, activeForegroundColor int32) Design {
	design.activeBackgroundColor = activeBackgroundColor
	design.activeForegroundColor = activeForegroundColor
	return design
}

func (design Design) SetInactiveStyle(inactiveBackgroundColor, inactiveForegroundColor int32) Design {
	design.inactiveBackgroundColor = inactiveBackgroundColor
	design.inactiveForegroundColor = inactiveForegroundColor
	return design
}

func (design Design) SetDisabledStyle(disabledBackgroundColor, disabledForegroundColor int32) Design {
	design.disabledBackgroundColor = disabledBackgroundColor
	design.disabledForegroundColor = disabledForegroundColor
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

func (design Design) GetActiveStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.activeBackgroundColor)).
		Foreground(tcell.NewHexColor(design.activeForegroundColor))

	return style
}

func (design Design) GetInactiveStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.inactiveBackgroundColor)).
		Foreground(tcell.NewHexColor(design.inactiveForegroundColor))

	return style
}

func (design Design) GetDisabledStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(design.disabledBackgroundColor)).
		Foreground(tcell.NewHexColor(design.disabledForegroundColor))

	return style
}

func (design Design) GetActiveStyleDecompose() (tcell.Color, tcell.Color) {
	activeBackgroundColor, activeForegroundColor, _ := design.GetActiveStyle().Decompose()
	return activeBackgroundColor, activeForegroundColor
}

func (design Design) GetInactiveStyleDecompose() (tcell.Color, tcell.Color) {
	inactiveBackgroundColor, inactiveForegroundColor, _ := design.GetInactiveStyle().Decompose()
	return inactiveBackgroundColor, inactiveForegroundColor
}

func (design Design) GetDisabledStyleDecompose() (tcell.Color, tcell.Color) {
	disabledBackgroundColor, disabledForegroundColor, _ := design.GetDisabledStyle().Decompose()
	return disabledBackgroundColor, disabledForegroundColor
}
