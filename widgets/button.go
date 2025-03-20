package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type button struct {
	label        string
	handleSelect func()
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

// NewButton creates and returns an instance of the button type.
//
// The button widget is used to create clickable buttons in Blaze. Properties
// such as the label and event handlers for selection and input can be customized.
// The widget aligns with the features and requirements of Blaze. Under the hood,
// it uses tview.Button to render the user interface of the button.
//
// # Properties
//
//	label        string
//	handleSelect func()
//	handleInput  func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*button
//
// # Usage
//
//	button := widgets.NewButton()
func NewButton() *button {
	return &button{
		label:        "",
		handleSelect: nil,
		handleInput:  nil,
	}
}

// SetLabel sets the label for the button type.
//
// A label is displayed on the button to indicate its purpose. Method chaining
// is supported.
//
// # Parameters
//
//	label string
//
// # Returns
//
//	*button
//
// # Usage
//
//	button := widgets.NewButton().SetLabel("")
func (button *button) SetLabel(label string) *button {
	button.label = label
	return button
}

// HandleSelect assigns a custom function to handle button selection.
//
// Selection behavior is customized with a function that runs when the button
// is clicked. Method chaining is supported.
//
// # Parameters
//
//	handleSelect func()
//
// # Returns
//
//	*button
//
// # Usage
//
//	button := widgets.
//	     NewButton().
//		 HandleSelect(
//		     func() {
//		     },
//		 )
func (button *button) HandleSelect(handleSelect func()) *button {
	button.handleSelect = handleSelect
	return button
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
//	*button
//
// # Usage
//
//	button := widgets.
//	     NewButton().
//		 HandleInput(
//		     func(event *tcell.EventKey) *tcell.EventKey {
//		         return event
//		     },
//		 )
func (button *button) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *button {
	button.handleInput = handleInput
	return button
}

// Render creates and returns a tview.Button with the configured properties.
//
// A tview.Button is built with the set properties. This function renders the
// button for display in Blaze.
//
// # Returns
//
//	*tview.Button
//
// # Usage
//
//	button := widgets.
//	     NewButton().
//	     SetLabel("").
//		 HandleSelect(
//		     func(event *tcell.EventKey) *tcell.EventKey {
//		         return event
//		     },
//		 ).
//		 HandleInput(
//		     func(event *tcell.EventKey) *tcell.EventKey {
//		         return event
//		     },
//		 ).
//		 Render()
func (button *button) Render() *tview.Button {
	element := tview.
		NewButton(button.label).
		SetSelectedFunc(button.handleSelect).
		SetStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetActivatedStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetDisabledStyle(
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		)

	element.
		SetBorder(true).
		SetInputCapture(button.handleInput).
		SetBackgroundColor(tcell.ColorBlack)

	return element
}
