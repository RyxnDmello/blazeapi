package area

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Area struct {
	text            string
	label           string
	placeholder     string
	rows            int
	columns         int
	maximum         int
	wrap            bool
	wordWrap        bool
	design          Design
	handleChange    func(value string, length int, firstChar, lastChar rune)
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewTextArea() *Area {
	return &Area{
		text:            "",
		label:           "",
		placeholder:     "",
		rows:            0,
		columns:         0,
		maximum:         0,
		wrap:            true,
		wordWrap:        true,
		design:          NewDesign(),
		handleChange:    nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (area *Area) SetText(text string) *Area {
	area.text = text
	return area
}

func (area *Area) SetLabel(label string) *Area {
	area.label = label
	return area
}

func (area *Area) SetPlaceholder(placeholder string) *Area {
	area.placeholder = placeholder
	return area
}

func (area *Area) SetSize(rows, columns int) *Area {
	area.rows = max(0, rows)
	area.columns = max(0, columns)
	return area
}

func (area *Area) SetMaximum(maximum int) *Area {
	area.maximum = max(0, maximum)
	return area
}

func (area *Area) SetWrap(wrap bool) *Area {
	area.wrap = wrap
	return area
}

func (area *Area) SetWordWrap(wordWrap bool) *Area {
	area.wordWrap = wordWrap
	return area
}

func (area *Area) SetDesign(design Design) *Area {
	area.design = design
	return area
}

func (area *Area) HandleChange(handleChange func(value string, length int, firstChar, lastChar rune)) *Area {
	area.handleChange = handleChange
	return area
}

func (area *Area) HandleSwitch(handleSwitch func()) *Area {
	area.handleSwitch = handleSwitch
	return area
}

func (area *Area) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Area {
	area.handleInput = handleInput
	return area
}

func (area *Area) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Area {
	area.handleDimension = handleDimension
	return area
}

func (area *Area) GetHeight() int {
	return max(1, area.rows)
}

func (area *Area) GetHandleChange() func(value string, length int, firstChar, lastChar rune) {
	return area.handleChange
}

func (area *Area) GetHandleSwitch() func() {
	return area.handleSwitch
}

func (area *Area) View() *tview.TextArea {
	widget := tview.
		NewTextArea().
		SetWrap(area.wrap).
		SetLabel(area.label).
		SetText(area.text, true).
		SetWordWrap(area.wordWrap).
		SetMaxLength(area.maximum).
		SetPlaceholder(area.placeholder).
		SetSize(area.rows, area.columns).
		SetTextStyle(area.design.GetTextStyle()).
		SetLabelStyle(area.design.GetLabelStyle()).
		SetPlaceholderStyle(area.design.GetPlaceholderStyle())

	widget.
		SetChangedFunc(func() {
			if area.handleChange != nil {
				text := widget.GetText()

				length, firstChar, lastChar := len(text), rune(0), rune(0)

				if length == 0 {
					firstChar, lastChar = rune(0), rune(0)
				}

				if length > 0 {
					lastChar = rune(text[length-1])
				}

				area.handleChange(text, length, firstChar, lastChar)
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(area.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if area.handleSwitch != nil && event.Key() == tcell.KeyTab {
				area.handleSwitch()
			}

			if area.handleInput != nil {
				return area.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range area.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), false)
			}

			if area.handleDimension == nil {
				return x, y, width, height
			}

			return area.handleDimension(screen, x, y, width, height)
		})

	return widget
}
