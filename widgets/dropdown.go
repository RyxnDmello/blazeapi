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

func NewDropdown() *dropdown {
	return &dropdown{
		options:      make([]string, 0),
		handleSelect: nil,
		handleInput:  nil,
	}
}

func (dropdown *dropdown) SetOptions(options []string) *dropdown {
	dropdown.options = options
	return dropdown
}

func (dropdown *dropdown) HandleSelect(handleSelect func(text string, index int)) *dropdown {
	dropdown.handleSelect = handleSelect
	return dropdown
}

func (dropdown *dropdown) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *dropdown {
	dropdown.handleInput = handleInput
	return dropdown
}

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
