package input

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Input struct {
	label           string
	placeholder     string
	design          Design
	autocomplete    Autocomplete
	handleClick     func()
	handleChange    func(value string, length int, firstChar, lastChar rune)
	handleAccept    func(value string, length int, firstChar, lastChar rune) bool
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewInput() *Input {
	return &Input{
		label:           "",
		placeholder:     "",
		design:          NewDesign(),
		autocomplete:    NewAutocomplete(),
		handleClick:     nil,
		handleChange:    nil,
		handleAccept:    nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (input *Input) SetLabel(label string) *Input {
	input.label = label
	return input
}

func (input *Input) SetPlaceholder(placeholder string) *Input {
	input.placeholder = placeholder
	return input
}

func (input *Input) SetDesign(design Design) *Input {
	input.design = design
	return input
}

func (input *Input) SetAutocomplete(autocomplete Autocomplete) *Input {
	input.autocomplete = autocomplete
	return input
}

func (input *Input) HandleClick(handleClick func()) *Input {
	input.handleClick = handleClick
	return input
}

func (input *Input) HandleChange(handleChange func(value string, length int, firstChar, lastChar rune)) *Input {
	input.handleChange = handleChange
	return input
}

func (input *Input) HandleAccept(handleAccept func(value string, length int, firstChar, lastChar rune) bool) *Input {
	input.handleAccept = handleAccept
	return input
}

func (input *Input) HandleSwitch(handleSwitch func()) *Input {
	input.handleSwitch = handleSwitch
	return input
}

func (input *Input) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Input {
	input.handleInput = handleInput
	return input
}

func (input *Input) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Input {
	input.handleDimension = handleDimension
	return input
}

func (input *Input) GetHandleChange() func(value string, length int, firstChar, lastChar rune) {
	return input.handleChange
}

func (input *Input) GetHandleSwitch() func() {
	return input.handleSwitch
}

func (input *Input) View() *tview.InputField {
	widget := tview.
		NewInputField().
		SetFieldWidth(0).
		SetLabel(input.label).
		SetPlaceholder(input.placeholder).
		SetLabelStyle(input.design.GetLabelStyle()).
		SetFieldStyle(input.design.GetFieldStyle()).
		SetPlaceholderStyle(input.design.GetPlaceholderStyle()).
		SetAutocompleteStyles(input.design.GetRootColor(), input.design.GetAutocompleteSelectedStyle(), input.design.GetAutocompleteUnselectedStyle())

	widget.
		SetChangedFunc(func(text string) {
			if input.handleChange != nil {
				length, firstChar, lastChar := len(text), rune(0), rune(0)

				if length == 0 {
					firstChar, lastChar = rune(0), rune(0)
				}

				if length > 0 {
					lastChar = rune(text[length-1])
				}

				input.handleChange(text, length, firstChar, lastChar)
			}
		}).
		SetAcceptanceFunc(func(text string, lastChar rune) bool {
			if input.handleAccept != nil {
				return input.handleAccept(text, len(text), rune(text[0]), lastChar)
			}

			return true
		}).
		SetAutocompleteFunc(func(text string) (options []string) {
			if input.autocomplete.GetEntries() == nil {
				return options
			}

			entries := input.autocomplete.GetEntries()

			for value, options := range entries {
				if !strings.Contains(text, value) {
					continue
				}

				length := 0
				available := make([]string, 0)
				leftPadding, rightPadding := input.autocomplete.GetPadding()

				for _, option := range options {
					if len(option) > length {
						length = len(option)
					}
				}

				for _, option := range options {
					left := "" + strings.Repeat(" ", leftPadding)
					right := "" + strings.Repeat(" ", rightPadding+length-len(option))
					available = append(available, left+option+right)
				}

				return available
			}

			return options
		}).
		SetAutocompletedFunc(func(option string, index, source int) bool {
			if input.autocomplete.GetHandleAutocomplete() != nil {
				return input.autocomplete.GetHandleAutocomplete()(option, len(option), source)
			}

			if source != tview.AutocompletedNavigate {
				value := strings.TrimSpace(option)
				widget.SetText(value)
			}

			return source == tview.AutocompletedEnter || source == tview.AutocompletedClick
		}).
		SetDoneFunc(func(key tcell.Key) {
			if input.handleClick != nil && key == tcell.KeyEnter {
				input.handleClick()
				return
			}

			if input.handleSwitch != nil && key == tcell.KeyTab {
				input.handleSwitch()
				return
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(input.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if input.handleSwitch != nil && event.Key() == tcell.KeyTab {
				input.handleSwitch()
			}

			if input.handleInput != nil {
				return input.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range input.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), false)
			}

			if input.handleDimension == nil {
				return x, y, width, height
			}

			return input.handleDimension(screen, x, y, width, height)
		})

	return widget
}
