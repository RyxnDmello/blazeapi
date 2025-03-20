package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type textArea struct {
	text        string
	label       string
	placeholder string
	rows        int
	columns     int
	wordWrap    bool
	handleInput func(event *tcell.EventKey) *tcell.EventKey
}

// NewTextArea creates and returns an instance of the textArea type.
//
// The textArea widget is used to create multiline input fields. Properties
// such as the text, label, placeholder, rows, columns, word wrap, and input
// handling can be completely customized. The widget aligns with the features
// and requirements of Blaze. Under the hood, the widget uses tview.TextArea
// to render the user interface of the text area.
//
// # Properties
//
//	text        string
//	label       string
//	placeholder string
//	rows        int
//	columns     int
//	wordWrap    bool
//	handleInput func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea()
func NewTextArea() *textArea {
	return &textArea{
		text:        "",
		label:       "",
		placeholder: "",
		rows:        0,
		columns:     0,
		wordWrap:    true,
		handleInput: nil,
	}
}

// SetText sets the text content of the textArea type.
//
// The text content can be set or updated dynamically with this function.
// Method chaining is supported.
//
// # Parameters
//
//	text string
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetText("")
func (textArea *textArea) SetText(text string) *textArea {
	textArea.text = text
	return textArea
}

// SetLabel sets the label for the textArea type.
//
// A label is placed above the text area to provide context or instructions.
// Method chaining is supported.
//
// # Parameters
//
//	label string
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetLabel("")
func (textArea *textArea) SetLabel(label string) *textArea {
	textArea.label = label
	return textArea
}

// SetPlaceholder sets the placeholder text for the textArea type.
//
// A placeholder hint is displayed when the text area is empty. Method chaining
// is supported.
//
// # Parameters
//
//	placeholder string
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetPlaceholder("")
func (textArea *textArea) SetPlaceholder(placeholder string) *textArea {
	textArea.placeholder = placeholder
	return textArea
}

// SetRows sets the number of visible rows in the textArea type.
//
// The vertical size is defined by the number of visible lines. If set to 0 or
// a negative value, tview auto-sizes the height of the text area. Method chaining
// is supported.
//
// # Parameters
//
//	rows int
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetRows(0)
func (textArea *textArea) SetRows(rows int) *textArea {
	textArea.rows = rows
	return textArea
}

// SetColumns sets the number of visible columns in the textArea type.
//
// The horizontal size is defined by the number of visible characters. If set
// to 0 or a negative value, tview auto-sizes the width of the text area. Method
// chaining is supported.
//
// # Parameters
//
//	columns int
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetColumns(0)
func (textArea *textArea) SetColumns(columns int) *textArea {
	textArea.columns = columns
	return textArea
}

// SetWordWrap enables or disables word wrapping in the textArea type.
//
// Word wrapping can be enabled for text to wrap or disabled for horizontal
// scrolling. Method chaining is supported.
//
// # Parameters
//
//	wordWrap bool
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.NewTextArea().SetWordWrap(true)
func (textArea *textArea) SetWordWrap(wordWrap bool) *textArea {
	textArea.wordWrap = wordWrap
	return textArea
}

// HandleInput assigns a custom function to process keyboard events.
//
// Keyboard behavior can be customized with a tcell EventKey handler. If nil,
// tview defaults are used. Method chaining is supported.
//
// # Parameters
//
//	handleInput func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*textArea
//
// # Usage
//
//	textArea := widgets.
//	     NewTextArea().
//		 HandleInput(
//		     func(event *tcell.EventKey) *tcell.EventKey {
//		         return event
//		     },
//		 )
func (textArea *textArea) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *textArea {
	textArea.handleInput = handleInput
	return textArea
}

// Render creates and returns a tview.TextArea with the configured properties.
//
// A tview.TextArea is built with the set properties. This function renders
// the text area for display in Blaze.
//
// # Returns
//
//	*tview.TextArea
//
// # Usage
//
//	textArea := widgets.
//	     NewTextArea().
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     ).
//	     Render()
func (textArea *textArea) Render() *tview.TextArea {
	area := tview.
		NewTextArea().
		SetLabel(textArea.label).
		SetText(textArea.text, true).
		SetWordWrap(textArea.wordWrap).
		SetPlaceholder(textArea.placeholder).
		SetSize(textArea.rows, textArea.columns).
		SetTextStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetLabelStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		).
		SetPlaceholderStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		)

	area.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetBackgroundColor(tcell.ColorBlack).
		SetInputCapture(textArea.handleInput)

	return area
}
