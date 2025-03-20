package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type input struct {
	label        string
	placeholder  string
	handleChange func(text string)
	handleAccept func(text string, lastChar rune) bool
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

// NewInput creates and returns an instance of the input type.
//
// The input widget is used to create single-line text input fields in Blaze.
// Properties such as the label, placeholder, and event handlers for change,
// acceptance, and input can be customized. The widget aligns with the features
// and requirements of Blaze. Under the hood, it uses tview.InputField to render
// the user interface.
//
// # Properties
//
//	label        string
//	placeholder  string
//	handleChange func(text string)
//	handleAccept func(text string, lastChar rune) bool
//	handleInput  func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*input
//
// # Usage
//
//	input := widgets.NewInput()
func NewInput() *input {
	return &input{
		label:        "",
		placeholder:  "",
		handleChange: nil,
		handleAccept: nil,
		handleInput:  nil,
	}
}

// SetLabel sets the label for the input type.
//
// A label is placed beside the input field to provide context or instructions.
// Method chaining is supported.
//
// # Parameters
//
//	label string
//
// # Returns
//
//	*input
//
// # Usage
//
//	input := widgets.NewInput().SetLabel("")
func (input *input) SetLabel(label string) *input {
	input.label = label
	return input
}

// SetPlaceholder sets the placeholder text for the input type.
//
// A placeholder hint is displayed when the input field is empty. Method chaining
// is supported.
//
// # Parameters
//
//	placeholder string
//
// # Returns
//
//	*input
//
// # Usage
//
//	input := widgets.NewInput().SetPlaceholder("")
func (input *input) SetPlaceholder(placeholder string) *input {
	input.placeholder = placeholder
	return input
}

// HandleChange assigns a custom function to handle text changes.
//
// Change behavior is customized with a function that runs when the input text
// is modified. Method chaining is supported.
//
// # Parameters
//
//	handleChange func(text string)
//
// # Returns
//
//	*input
//
// # Usage
//
//	input := widgets.
//	     NewInput().
//	     HandleChange(
//	         func(text string) {
//	         },
//	     )
func (input *input) HandleChange(handleChange func(text string)) *input {
	input.handleChange = handleChange
	return input
}

// HandleAcceptance assigns a custom function to handle input acceptance.
//
// Acceptance behavior is customized with a function that validates the input
// when submitted. Method chaining is supported.
//
// # Parameters
//
//	handleAccept func(text string, lastChar rune) bool
//
// # Returns
//
//	*input
//
// # Usage
//
//	input := widgets.
//	     NewInput().
//	     HandleAcceptance(
//	         func(text string, lastChar rune) bool {
//	             return true
//	         },
//	     )
func (input *input) HandleAcceptance(handleAccept func(text string, lastChar rune) bool) *input {
	input.handleAccept = handleAccept
	return input
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
//	*input
//
// # Usage
//
//	input := widgets.
//	     NewInput().
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     )
func (input *input) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *input {
	input.handleInput = handleInput
	return input
}

// Render creates and returns a tview.InputField with the configured properties.
//
// A tview.InputField is built with the set properties. This function renders
// the input field for display in Blaze.
//
// # Returns
//
//	*tview.InputField
//
// # Usage
//
//	input := widgets.
//	     NewInput().
//	     SetLabel("").
//	     SetPlaceholder("").
//	     HandleChange(
//	         func(text string) {
//	         },
//	     ).
//	     HandleAcceptance(
//	         func(text string, lastChar rune) bool {
//	             return true
//	         },
//	     ).
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     ).
//	     Render()
func (input *input) Render() (field *tview.InputField) {
	field = tview.
		NewInputField().
		SetLabel(input.label).
		SetPlaceholder(input.placeholder).
		SetChangedFunc(input.handleChange).
		SetAcceptanceFunc(input.handleAccept).
		SetFieldStyle(
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

	field.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetInputCapture(input.handleInput).
		SetBackgroundColor(tcell.ColorBlack)

	return field
}
