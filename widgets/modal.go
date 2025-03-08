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

func (modal *modal) SetTitle(title string) *modal {
	modal.title = title
	return modal
}

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

func (modal *modal) SetDimension(width, height int) *modal {
	modal.dimension.width = width
	modal.dimension.height = height
	return modal
}

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

func flow(elements []Element, direction int) (flow *tview.Flex) {
	flow = tview.
		NewFlex().
		SetDirection(direction)

	for _, element := range elements {
		flow.AddItem(element.entity, 0, 1, element.active)
	}

	return flow
}
