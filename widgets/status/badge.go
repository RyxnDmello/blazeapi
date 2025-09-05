package status

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Badge struct {
	icon            *rune
	value           string
	leftPadding     int
	rightPadding    int
	leftGlyphs      []Glyph
	rightGlyphs     []Glyph
	backgroundColor int32
	foregroundColor int32
}

func NewBadge() *Badge {
	return &Badge{
		icon:            nil,
		value:           "",
		leftPadding:     1,
		rightPadding:    1,
		leftGlyphs:      make([]Glyph, 0),
		rightGlyphs:     make([]Glyph, 0),
		backgroundColor: tcell.ColorDefault.Hex(),
		foregroundColor: tcell.ColorDefault.Hex(),
	}
}

func (badge *Badge) SetIcon(icon rune) *Badge {
	badge.icon = &icon
	return badge
}

func (badge *Badge) SetValue(value string) *Badge {
	badge.value = value
	return badge
}

func (badge *Badge) SetPadding(leftPadding, rightPadding int) *Badge {
	badge.leftPadding = max(0, leftPadding)
	badge.rightPadding = max(0, rightPadding)
	return badge
}

func (badge *Badge) AddLeftGlyph(leftGlyph Glyph) *Badge {
	badge.leftGlyphs = slices.Concat([]Glyph{leftGlyph}, badge.leftGlyphs)
	return badge
}

func (badge *Badge) AddRightGlyph(rightGlyph Glyph) *Badge {
	badge.rightGlyphs = slices.Concat([]Glyph{rightGlyph}, badge.rightGlyphs)
	return badge
}

func (badge *Badge) SetStyle(backgroundColor, foregroundColor int32) *Badge {
	badge.backgroundColor = backgroundColor
	badge.foregroundColor = foregroundColor
	return badge
}

func (badge *Badge) GetDisplay() string {
	name := badge.value

	if badge.icon != nil {
		name = string(*badge.icon) + name
	}

	return name
}

func (badge *Badge) GetWidth() int {
	width := badge.leftPadding + len(badge.leftGlyphs) + len(badge.value) + len(badge.rightGlyphs) + badge.rightPadding

	if badge.icon != nil {
		return width + 1
	}

	return width
}

func (badge *Badge) GetRootColor() tcell.Color {
	return tcell.NewHexColor(badge.backgroundColor)
}

func (badge *Badge) GetStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(badge.backgroundColor)).
		Foreground(tcell.NewHexColor(badge.foregroundColor))

	return style
}

func (badge *Badge) GetLayout() (*tview.TextView, int) {
	width := badge.
		GetWidth()

	widget := tview.
		NewTextView().
		SetWrap(false).
		SetWordWrap(false).
		SetDynamicColors(true).
		SetText(badge.GetDisplay()).
		SetTextStyle(badge.GetStyle()).
		SetTextAlign(tview.AlignCenter)

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(badge.GetRootColor())

	widget.
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for index, glyph := range badge.leftGlyphs {
				screen.SetContent(x+index, y, glyph.symbol, nil, glyph.GetStyle())
			}

			for index, glyph := range badge.rightGlyphs {
				screen.SetContent(x+width-1-index, y, glyph.symbol, nil, glyph.GetStyle())
			}

			return x + len(badge.leftGlyphs), y, width - len(badge.rightGlyphs), height
		})

	return widget, width
}
