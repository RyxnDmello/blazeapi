package button

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Button struct {
	label           string
	disabled        bool
	design          Design
	handleClick     func()
	handleExit      func(key tcell.Key)
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewButton() *Button {
	return &Button{
		label:           "",
		disabled:        false,
		design:          NewDesign(),
		handleClick:     nil,
		handleExit:      nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (button *Button) SetLabel(label string) *Button {
	button.label = label
	return button
}

func (button *Button) SetDisabled(disabled bool) *Button {
	button.disabled = disabled
	return button
}

func (button *Button) SetDesign(design Design) *Button {
	button.design = design
	return button
}

func (button *Button) HandleClick(handleClick func()) *Button {
	button.handleClick = handleClick
	return button
}

func (button *Button) HandleExit(handleExit func(key tcell.Key)) *Button {
	button.handleExit = handleExit
	return button
}

func (button *Button) HandleSwitch(handleSwitch func()) *Button {
	button.handleSwitch = handleSwitch
	return button
}

func (button *Button) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Button {
	button.handleInput = handleInput
	return button
}

func (button *Button) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Button {
	button.handleDimension = handleDimension
	return button
}

func (button *Button) GetHandleClick() func() {
	return button.handleClick
}

func (button *Button) GetHandleSwitch() func() {
	return button.handleSwitch
}

func (button *Button) View() *tview.Button {
	widget := tview.
		NewButton(button.label).
		SetLabel(button.label).
		SetDisabled(button.disabled).
		SetStyle(button.design.GetInactiveStyle()).
		SetActivatedStyle(button.design.GetActiveStyle()).
		SetDisabledStyle(button.design.GetDisabledStyle())

	widget.
		SetSelectedFunc(func() {
			if button.handleClick != nil {
				button.handleClick()
			}
		}).
		SetExitFunc(func(key tcell.Key) {
			if button.handleSwitch != nil && key == tcell.KeyTAB {
				button.handleSwitch()
			}

			if button.handleExit != nil {
				button.handleExit(key)
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(button.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if button.handleSwitch != nil && event.Key() == tcell.KeyTab {
				button.handleSwitch()
			}

			if button.handleInput != nil {
				return button.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range button.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), widget.IsDisabled())
			}

			if button.handleDimension == nil {
				return x, y, width, height
			}

			return button.handleDimension(screen, x, y, width, height)
		})

	return widget
}
