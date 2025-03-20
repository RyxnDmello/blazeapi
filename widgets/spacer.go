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

// NewSpacer creates and returns an instance of the spacer type.
//
// The spacer widget is used to create empty space or separators in Blaze.
// Properties such as the border, border color, and background color can be
// customized. The widget aligns with the features and requirements of Blaze.
// Under the hood, it uses tview.Box to render the user interface.
//
// # Properties
//
//	border          bool
//	borderColor     tcell.Color
//	backgroundColor tcell.Color
//
// # Returns
//
//	*spacer
//
// # Usage
//
//	spacer := widgets.NewSpacer()
func NewSpacer() *spacer {
	return &spacer{
		border:          true,
		borderColor:     tcell.ColorBlack,
		backgroundColor: tcell.ColorBlack,
	}
}

// SetBorder sets the border visibility for the spacer type.
//
// The border can be enabled or disabled to define the spacer’s edges. Method
// chaining is supported.
//
// # Parameters
//
//	border bool
//
// # Returns
//
//	*spacer
//
// # Usage
//
//	spacer := widgets.NewSpacer().SetBorder(true)
func (spacer *spacer) SetBorder(border bool) *spacer {
	spacer.border = border
	return spacer
}

// SetBorderColor sets the border color for the spacer type.
//
// The border color can be customized to change the spacer’s appearance. Method
// chaining is supported.
//
// # Parameters
//
//	borderColor tcell.Color
//
// # Returns
//
//	*spacer
//
// # Usage
//
//	spacer := widgets.NewSpacer().SetBorderColor(tcell.ColorBlack)
func (spacer *spacer) SetBorderColor(borderColor tcell.Color) *spacer {
	spacer.borderColor = borderColor
	return spacer
}

// SetBackgroundColor sets the background color for the spacer type.
//
// The background color can be customized to fill the spacer’s area. Method
// chaining is supported.
//
// # Parameters
//
//	backgroundColor tcell.Color
//
// # Returns
//
//	*spacer
//
// # Usage
//
//	spacer := widgets.NewSpacer().SetBackgroundColor(tcell.ColorBlack)
func (spacer *spacer) SetBackgroundColor(backgroundColor tcell.Color) *spacer {
	spacer.backgroundColor = backgroundColor
	return spacer
}

// Render creates and returns a tview.Box with the configured properties.
//
// A tview.Box is built with the set properties. This function renders the
// spacer for display in Blaze.
//
// # Returns
//
//	*tview.Box
//
// # Usage
//
//	spacer := widgets.
//	     NewSpacer().
//	     SetBorder(true).
//	     SetBorderColor(tcell.ColorBlack).
//	     SetBackgroundColor(tcell.ColorBlack).
//	     Render()
func (spacer *spacer) Render() *tview.Box {
	container := tview.
		NewBox().
		SetBorder(spacer.border).
		SetBorderColor(spacer.borderColor).
		SetBackgroundColor(spacer.backgroundColor)

	return container
}
