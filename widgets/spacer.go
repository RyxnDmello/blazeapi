package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type spacer struct {
	border          bool
	borderColor     tcell.Color
	backgroundColor tcell.Color
}

func NewSpacer() *spacer {
	return &spacer{
		border:          true,
		borderColor:     tcell.ColorBlack,
		backgroundColor: tcell.ColorBlack,
	}
}

func (spacer *spacer) SetBorder(border bool) *spacer {
	spacer.border = border
	return spacer
}

func (spacer *spacer) SetBorderColor(borderColor tcell.Color) *spacer {
	spacer.borderColor = borderColor
	return spacer
}

func (spacer *spacer) SetBackgroundColor(backgroundColor tcell.Color) *spacer {
	spacer.backgroundColor = backgroundColor
	return spacer
}

func (spacer *spacer) Render() *tview.Box {
	container := tview.
		NewBox().
		SetBorder(spacer.border).
		SetBorderColor(spacer.borderColor).
		SetBackgroundColor(spacer.borderColor)

	return container
}
