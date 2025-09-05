package status

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Status struct {
	sections  []*Section
	rootColor int32
}

func NewStatus() *Status {
	return &Status{
		sections:  make([]*Section, 0),
		rootColor: tcell.ColorDefault.Hex(),
	}
}

func (status *Status) AddSection(section *Section) *Status {
	status.sections = append(status.sections, section)
	return status
}

func (status *Status) SetRootColor(rootColor int32) *Status {
	status.rootColor = rootColor
	return status
}

func (status *Status) GetSpacer() *tview.Box {
	box := tview.
		NewBox().
		SetBackgroundColor(status.GetRootColor())

	return box
}

func (status *Status) GetRootColor() tcell.Color {
	return tcell.NewHexColor(status.rootColor)
}

func (status *Status) View() *tview.Flex {
	spacer := status.
		GetSpacer()

	widget := tview.
		NewFlex().
		SetDirection(tview.FlexColumn)

	if len(status.sections) == 0 {
		widget.
			AddItem(spacer, 0, 1, false)

		return widget
	}

	for index, section := range status.sections {
		layout, width := section.
			GetLayout()

		widget.
			AddItem(layout, width, 1, false)

		if index == 0 || index < len(status.sections)-1 {
			widget.
				AddItem(spacer, 0, 1, false)
		}
	}

	return widget
}
