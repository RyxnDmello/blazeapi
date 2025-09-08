package modal

import (
	"github.com/ryxndmello/flame/widgets/button"
	"github.com/ryxndmello/flame/widgets/glyph"
	"github.com/ryxndmello/flame/widgets/input"
	"github.com/ryxndmello/flame/widgets/message"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	InputModalInput         = 0
	InputModalSuccessButton = 1
	InputModalFailureButton = 2
)

type InputModal struct {
	message             *message.Message
	input               *input.Input
	success             *button.Button
	failure             *button.Button
	handleDefaultClick  func(active int)
	handleDefaultSwitch func(active int) (*tview.Application, int)
	handleDefaultChange func(value string, length int, firstChar, lastChar rune)
}

func NewInputModal() *InputModal {
	return &InputModal{
		message:             nil,
		input:               nil,
		success:             nil,
		failure:             nil,
		handleDefaultClick:  nil,
		handleDefaultSwitch: nil,
		handleDefaultChange: nil,
	}
}

func (modal *InputModal) SetMessage(message *message.Message) *InputModal {
	modal.message = message
	return modal
}

func (modal *InputModal) SetDefaultMessage(text string) *InputModal {
	design := message.
		NewDesign().
		SetRootColor(0x11111b).
		SetStyle(0x11111b, 0xcdd6f4)

	modal.message = message.
		NewMessage().
		SetText(text).
		SetDesign(design).
		SetAlignment(tview.AlignCenter)

	return modal
}

func (modal *InputModal) SetInput(input *input.Input) *InputModal {
	modal.input = input
	return modal
}

func (modal *InputModal) SetDefaultInput(label, placeholder string) *InputModal {
	design := input.
		NewDesign().
		SetRootColor(0x0e0e15).
		SetFieldStyle(0x0e0e15, 0xcdd6f4).
		SetLabelStyle(0x0e0e15, 0xcdd6f4).
		SetPlaceholderStyle(0x0e0e15, 0xcdd6f4)

	modal.input = input.
		NewInput().
		SetLabel(label).
		SetDesign(design).
		SetPlaceholder(placeholder)

	return modal
}

func (modal *InputModal) SetSuccessButton(success *button.Button) *InputModal {
	modal.success = success
	return modal
}

func (modal *InputModal) SetDefaultSuccessButton(label string) *InputModal {
	startGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x11111b, 0x82b67d).
		SetFocusedStyle(0x11111b, 0xa6e3a1).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	endGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x11111b, 0x82b67d).
		SetFocusedStyle(0x11111b, 0xa6e3a1).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	design := button.
		NewDesign().
		SetRootColor(0x11111b).
		SetActiveStyle(0xa6e3a1, 0x1e1e2e).
		SetInactiveStyle(0x82b67d, 0x1e1e2e).
		AddGlyph(startGlyph).
		AddGlyph(endGlyph)

	modal.success = button.
		NewButton().
		SetLabel(label).
		SetDesign(design).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 2, height
		})

	return modal
}

func (modal *InputModal) SetFailureButton(failure *button.Button) *InputModal {
	modal.failure = failure
	return modal
}

func (modal *InputModal) SetDefaultFailureButton(label string) *InputModal {
	startGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0ba').
		SetStyle(0x11111b, 0xbf6078).
		SetFocusedStyle(0x11111b, 0xf38ba8).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX, elementY
		})

	endGlyph := glyph.
		NewGlyph().
		SetSymbol('\ue0bc').
		SetStyle(0x11111b, 0xbf6078).
		SetFocusedStyle(0x11111b, 0xf38ba8).
		HandlePosition(func(elementX, elementY, elementWidth, elementHeight int) (int, int) {
			return elementX + elementWidth - 1, elementY
		})

	design := button.
		NewDesign().
		SetRootColor(0x11111b).
		SetActiveStyle(0xf38ba8, 0x11111b).
		SetInactiveStyle(0xbf6078, 0x11111b).
		AddGlyph(startGlyph).
		AddGlyph(endGlyph)

	modal.failure = button.
		NewButton().
		SetLabel(label).
		SetDesign(design).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 1, y, width - 2, height
		})

	return modal
}

func (modal *InputModal) HandleDefaultClick(handleDefaultClick func(active int)) *InputModal {
	modal.handleDefaultClick = handleDefaultClick
	return modal
}

func (modal *InputModal) HandleDefaultSwitch(handleDefaultSwitch func(active int) (*tview.Application, int)) *InputModal {
	modal.handleDefaultSwitch = handleDefaultSwitch
	return modal
}

func (modal *InputModal) HandleDefaultChange(handleDefaultChange func(value string, length int, firstChar, lastChar rune)) *InputModal {
	modal.handleDefaultChange = handleDefaultChange
	return modal
}

func (modal *InputModal) GetLayout(rowGap, columnGap int, spacer *tview.Box) (*tview.Flex, int) {
	var message *tview.TextView
	var input *tview.InputField
	var success *tview.Button
	var failure *tview.Button

	var height int = 3
	var layout *tview.Flex
	var buttons *tview.Flex

	if modal.message == nil {
		modal.
			SetDefaultMessage("Message")
	}

	message = modal.message.
		View()

	if modal.input == nil {
		modal.
			SetDefaultInput("", "Enter Message")
	}

	handleChange := modal.input.
		GetHandleChange()

	handleSwitch := modal.input.
		GetHandleSwitch()

	input = modal.input.
		HandleChange(func(value string, length int, firstChar, lastChar rune) {
			if modal.handleDefaultChange != nil {
				modal.handleDefaultChange(value, length, firstChar, lastChar)
				return
			}

			if handleChange != nil {
				handleChange(value, length, firstChar, lastChar)
			}
		}).
		HandleSwitch(func() {
			if modal.handleDefaultSwitch != nil {
				app, active := modal.handleDefaultSwitch(InputModalInput)

				if active == InputModalSuccessButton && success != nil {
					app.SetFocus(success)
				}

				if active == InputModalFailureButton && failure != nil {
					app.SetFocus(failure)
				}

				return
			}

			if handleSwitch != nil {
				handleSwitch()
			}
		}).
		View()

	if modal.success == nil && modal.failure == nil {
		modal.
			SetDefaultSuccessButton("Success").
			SetDefaultFailureButton("Failure")
	}

	if modal.success != nil {
		handleClick := modal.success.
			GetHandleClick()

		handleSwitch := modal.success.
			GetHandleSwitch()

		success = modal.success.
			HandleClick(func() {
				if modal.handleDefaultClick != nil {
					modal.handleDefaultClick(InputModalSuccessButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(InputModalSuccessButton)

					if index == InputModalInput {
						app.SetFocus(input)
					}

					if index == InputModalFailureButton && failure != nil {
						app.SetFocus(failure)
					}

					return
				}

				if handleSwitch != nil {
					handleSwitch()
				}
			}).
			View()
	}

	if modal.failure != nil {
		handleClick := modal.failure.
			GetHandleClick()

		handleSwitch := modal.failure.
			GetHandleSwitch()

		failure = modal.failure.
			HandleClick(func() {
				if modal.handleDefaultClick != nil {
					modal.handleDefaultClick(InputModalFailureButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(InputModalFailureButton)

					if index == InputModalInput {
						app.SetFocus(input)
					}

					if index == InputModalSuccessButton && success != nil {
						app.SetFocus(success)
					}

					return
				}

				if handleSwitch != nil {
					handleSwitch()
				}
			}).
			View()
	}

	buttons = tview.
		NewFlex().
		SetDirection(tview.FlexColumn)

	if success != nil && failure == nil {
		buttons.
			AddItem(success, 0, 1, true)
	}

	if failure != nil && success == nil {
		buttons.
			AddItem(failure, 0, 1, true)
	}

	if success != nil && failure != nil {
		buttons.
			AddItem(success, 0, 1, true)

		if columnGap > 0 {
			buttons.
				AddItem(spacer, columnGap, 1, false)
		}

		buttons.
			AddItem(failure, 0, 1, true)
	}

	layout = tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(message, 1, 1, false)

	if rowGap > 0 {
		layout.
			AddItem(spacer, rowGap, 1, false)

		height += rowGap
	}

	layout.
		AddItem(input, 1, 1, true)

	if rowGap > 0 {
		layout.
			AddItem(spacer, rowGap, 1, false)

		height += rowGap
	}

	layout.
		AddItem(buttons, 1, 1, true)

	return layout, height
}
