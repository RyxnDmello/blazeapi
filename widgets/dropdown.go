package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type dropdown struct {
	options      []string
	handleSelect func(text string, index int)
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

// NewDropdown creates and returns an instance of the dropdown type.
//
// The dropdown widget is used to create selectable option lists in Blaze.
// Properties such as the options and event handlers for selection and input
// can be customized. The widget aligns with the features and requirements
// of Blaze. Under the hood, it uses tview.DropDown to render the user interface.
//
// # Properties
//
//	options      []string
//	handleSelect func(text string, index int)
//	handleInput  func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*dropdown
//
// # Usage
//
//	dropdown := widgets.NewDropdown()
func NewDropdown() *dropdown {
	return &dropdown{
		options:      make([]string, 0),
		handleSelect: nil,
		handleInput:  nil,
	}
}

// SetOptions sets the options for the dropdown type.
//
// A list of selectable options is defined for the dropdown. Method chaining
// is supported.
//
// # Parameters
//
//	options []string
//
// # Returns
//
//	*dropdown
//
// # Usage
//
//	dropdown := widgets.NewDropdown().SetOptions([]string{})
func (dropdown *dropdown) SetOptions(options []string) *dropdown {
	dropdown.options = options
	return dropdown
}

// HandleSelect assigns a custom function to handle dropdown selection.
//
// Selection behavior is customized with a function that runs when an option
// is chosen. Method chaining is supported.
//
// # Parameters
//
//	handleSelect func(text string, index int)
//
// # Returns
//
//	*dropdown
//
// # Usage
//
//	dropdown := widgets.
//	     NewDropdown().
//		 HandleSelect(
//		     func(text string, index int) {
//		     },
//		 )
func (dropdown *dropdown) HandleSelect(handleSelect func(text string, index int)) *dropdown {
	dropdown.handleSelect = handleSelect
	return dropdown
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
//	*dropdown
//
// # Usage
//
//	dropdown := widgets.
//	     NewDropdown().
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     )
func (dropdown *dropdown) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *dropdown {
	dropdown.handleInput = handleInput
	return dropdown
}

// Render creates and returns a tview.DropDown with the configured properties.
//
// A tview.DropDown is built with the set properties. This function renders
// the dropdown for display in Blaze.
//
// # Returns
//
//	*tview.DropDown
//
// # Usage
//
//	dropdown := widgets.
//	     NewDropdown().
//	     SetOptions([]string{}).
//	     HandleSelect(
//	         func(text string, index int) {
//	         },
//	     ).
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     ).
//	     Render()
func (dropdown *dropdown) Render() *tview.DropDown {
	element := tview.
		NewDropDown().
		SetCurrentOption(0).
		SetOptions(dropdown.options, nil).
		SetSelectedFunc(dropdown.handleSelect).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetTextOptions("  ", "   ", "  ", "", "  "+dropdown.options[0]).
		SetListStyles(
			tcell.StyleDefault.
				Bold(false).
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
			tcell.StyleDefault.
				Bold(true).
				Background(tcell.ColorNavy).
				Foreground(tcell.ColorWhite),
		)

	element.
		SetBorder(true).
		SetBackgroundColor(tcell.ColorBlack).
		SetInputCapture(dropdown.handleInput)

	return element
}
