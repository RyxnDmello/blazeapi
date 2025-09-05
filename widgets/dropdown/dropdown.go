package dropdown

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Dropdown struct {
	index           int
	options         []string
	prefix          string
	suffix          string
	listPrefix      string
	listSuffix      string
	design          Design
	handleSelect    func(option string, length, index int)
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewDropdown() *Dropdown {
	return &Dropdown{
		index:           0,
		options:         make([]string, 0),
		prefix:          "",
		suffix:          "",
		listPrefix:      "",
		listSuffix:      "",
		design:          NewDesign(),
		handleSelect:    nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (dropdown *Dropdown) SetIndex(index int) *Dropdown {
	dropdown.index = index
	return dropdown
}

func (dropdown *Dropdown) SetOptions(options []string) *Dropdown {
	dropdown.options = options
	return dropdown
}

func (dropdown *Dropdown) SetPrefix(prefix string) *Dropdown {
	dropdown.prefix = prefix
	return dropdown
}

func (dropdown *Dropdown) SetSuffix(suffix string) *Dropdown {
	dropdown.suffix = suffix
	return dropdown
}

func (dropdown *Dropdown) SetListPrefix(listPrefix string) *Dropdown {
	dropdown.listPrefix = listPrefix
	return dropdown
}

func (dropdown *Dropdown) SetListSuffix(listSuffix string) *Dropdown {
	dropdown.listSuffix = listSuffix
	return dropdown
}

func (dropdown *Dropdown) SetDesign(design Design) *Dropdown {
	dropdown.design = design
	return dropdown
}

func (dropdown *Dropdown) HandleSelect(handleSelect func(option string, length, index int)) *Dropdown {
	dropdown.handleSelect = handleSelect
	return dropdown
}

func (dropdown *Dropdown) HandleSwitch(handleSwitch func()) *Dropdown {
	dropdown.handleSwitch = handleSwitch
	return dropdown
}

func (dropdown *Dropdown) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Dropdown {
	dropdown.handleInput = handleInput
	return dropdown
}

func (dropdown *Dropdown) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Dropdown {
	dropdown.handleDimension = handleDimension
	return dropdown
}

func (dropdown *Dropdown) GetIndex() int {
	if dropdown.index < 0 {
		return 0
	}

	if dropdown.index > len(dropdown.options)-1 {
		return len(dropdown.options) - 1
	}

	return dropdown.index
}

func (dropdown *Dropdown) View() *tview.DropDown {
	index := dropdown.
		GetIndex()

	widget := tview.
		NewDropDown().
		SetFieldWidth(0).
		SetCurrentOption(index).
		SetOptions(dropdown.options, nil).
		SetTextOptions(dropdown.listPrefix, dropdown.listSuffix, dropdown.prefix, dropdown.suffix, dropdown.prefix+dropdown.options[index]+dropdown.suffix).
		SetLabelStyle(dropdown.design.GetLabelStyle()).
		SetFieldStyle(dropdown.design.GetFieldStyle()).
		SetFocusedStyle(dropdown.design.GetFocusedStyle()).
		SetDisabledStyle(dropdown.design.GetDisabledStyle()).
		SetListStyles(dropdown.design.GetListUnselectedStyle(), dropdown.design.GetListSelectedStyle())

	widget.
		SetSelectedFunc(func(text string, index int) {
			if dropdown.handleSelect != nil {
				dropdown.handleSelect(text, len(text), index)
			}
		}).
		SetDoneFunc(func(key tcell.Key) {
			if dropdown.handleSwitch != nil && key == tcell.KeyTAB {
				dropdown.handleSwitch()
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(dropdown.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if dropdown.handleSwitch != nil && event.Key() == tcell.KeyTab {
				dropdown.handleSwitch()
			}

			if dropdown.handleInput != nil {
				return dropdown.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range dropdown.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), false)
			}

			if dropdown.handleDimension == nil {
				return x, y, width, height
			}

			return dropdown.handleDimension(screen, x, y, width, height)
		})

	return widget
}
