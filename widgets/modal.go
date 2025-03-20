package widgets

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

type modal struct {
	title     string
	inputs    []Element
	buttons   []Element
	dimension Dimension
}

// NewModal creates and returns an instance of the modal type.
//
// The modal widget is used to create popup dialogs in Blaze. Properties such
// as the title, inputs, buttons, and dimensions can be customized. The widget
// aligns with the features and requirements of Blaze. Under the hood, it uses
// tview.Flex to render the user interface.
//
// # Properties
//
//	title     string
//	inputs    []Element
//	buttons   []Element
//	dimension Dimension
//
// # Returns
//
//	*modal
//
// # Usage
//
//	modal := widgets.NewModal()
func NewModal() *modal {
	return &modal{
		title:   "",
		inputs:  make([]Element, 0),
		buttons: make([]Element, 0),
		dimension: Dimension{
			width:  0,
			height: 0,
		},
	}
}

// SetTitle sets the title of the modal type.
//
// The title is displayed at the top of the modal dialog. Method chaining is
// supported.
//
// # Parameters
//
//	title string
//
// # Returns
//
//	*modal
//
// # Usage
//
//	modal := widgets.NewModal().SetTitle("")
func (modal *modal) SetTitle(title string) *modal {
	modal.title = title
	return modal
}

// AddInput adds an input element to the modal type.
//
// An input element is appended to the modal’s input list with an active state.
// Method chaining is supported.
//
// # Parameters
//
//	input  tview.Primitive
//	active bool
//
// # Returns
//
//	*modal
//
// # Usage
//
//	modal := widgets.NewModal().AddInput(nil, false)
func (modal *modal) AddInput(input tview.Primitive, active bool) *modal {
	modal.inputs = append(
		modal.inputs,
		Element{
			entity: input,
			active: active,
		},
	)

	return modal
}

// AddButton adds a button element to the modal type.
//
// A button element is appended to the modal’s button list with an active state.
// Method chaining is supported.
//
// # Parameters
//
//	button tview.Primitive
//	active bool
//
// # Returns
//
//	*modal
//
// # Usage
//
//	modal := widgets.NewModal().AddButton(nil, false)
func (modal *modal) AddButton(button tview.Primitive, active bool) *modal {
	modal.buttons = append(
		modal.buttons,
		Element{
			entity: button,
			active: active,
		},
	)

	return modal
}

// SetDimension sets the dimensions of the modal type.
//
// The width and height are updated to the specified values. Method chaining
// is supported.
//
// # Parameters
//
//	width  int
//	height int
//
// # Returns
//
//	*modal
//
// # Usage
//
//	modal := widgets.NewModal().SetDimension(0, 0)
func (modal *modal) SetDimension(width, height int) *modal {
	modal.dimension.width = width
	modal.dimension.height = height
	return modal
}

// Render creates and returns a tview.Flex with the configured properties.
//
// A tview.Flex is built with the set properties. This function renders the
// modal dialog for display in Blaze.
//
// # Returns
//
//	*tview.Flex
//
// # Usage
//
//	modal := widgets.
//	     NewModal().
//	     SetTitle("").
//	     AddInput(nil, false).
//	     AddButton(nil, false).
//	     SetDimension(0, 0).
//	     Render()
func (modal *modal) Render() *tview.Flex {
	if len(modal.inputs) == 0 {
		panic("Modal Must Contain An Input")
	}

	if len(modal.buttons) == 0 {
		panic("Modal Must Contain A Button")
	}

	structure := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(NewSpacer().Render(), 1, 1, false).
		AddItem(flow(modal.inputs, tview.FlexRow), 0, 1, true).
		AddItem(flow(modal.buttons, tview.FlexColumn), 3, 1, true).
		AddItem(NewSpacer().Render(), 1, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(NewSpacer().Render(), 2, 1, false).
		AddItem(structure, 0, 1, true).
		AddItem(NewSpacer().Render(), 2, 1, false)

	layout.SetBorder(true).
		SetTitle(fmt.Sprintf(" %s %s[red][ X ]", modal.title, strings.Repeat("═", modal.dimension.width-len(modal.title)-13))).
		SetTitleAlign(tview.AlignCenter)

	alignment := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(NewSpacer().Render(), 1, 1, false).
		AddItem(layout, modal.dimension.height, 1, true).
		AddItem(NewSpacer().Render(), 0, 0, false).
		AddItem(nil, 0, 1, false)

	popup := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(alignment, modal.dimension.width, 1, true).
		AddItem(nil, 0, 1, false)

	return popup
}

// flow creates and returns a tview.Flex with the specified elements and direction.
//
// A tview.Flex is built to arrange elements in a row or column. This function
// is a helper for laying out modal components.
//
// # Parameters
//
//	elements  []Element
//	direction int
//
// # Returns
//
//	*tview.Flex
//
// # Usage
//
//	flow := widgets.flow([]Element{}, tview.FlexRow)
func flow(elements []Element, direction int) (flow *tview.Flex) {
	flow = tview.
		NewFlex().
		SetDirection(direction)

	for _, element := range elements {
		flow.AddItem(element.entity, 0, 1, element.active)
	}

	return flow
}
