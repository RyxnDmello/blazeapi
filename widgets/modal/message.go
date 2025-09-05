package modal

import (
	"blazeapi/widgets/button"
	"blazeapi/widgets/glyph"
	"blazeapi/widgets/message"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	MessageModalSuccessButton = 0
	MessageModalFailureButton = 1
)

type MessageModal struct {
	message             *message.Message
	success             *button.Button
	failure             *button.Button
	handleDefaultClick  func(active int)
	handleDefaultSwitch func(active int) (*tview.Application, int)
}

func NewMessageModal() *MessageModal {
	return &MessageModal{
		message:             nil,
		success:             nil,
		failure:             nil,
		handleDefaultClick:  nil,
		handleDefaultSwitch: nil,
	}
}

func (modal *MessageModal) SetMessage(message *message.Message) *MessageModal {
	modal.message = message
	return modal
}

func (modal *MessageModal) SetDefaultMessage(text string) *MessageModal {
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

func (modal *MessageModal) SetSuccessButton(success *button.Button) *MessageModal {
	modal.success = success
	return modal
}

func (modal *MessageModal) SetDefaultSuccessButton(label string) *MessageModal {
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

func (modal *MessageModal) SetFailureButton(failure *button.Button) *MessageModal {
	modal.failure = failure
	return modal
}

func (modal *MessageModal) SetDefaultFailureButton(label string) *MessageModal {
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

func (modal *MessageModal) HandleDefaultClick(handleDefaultClick func(active int)) *MessageModal {
	modal.handleDefaultClick = handleDefaultClick
	return modal
}

func (modal *MessageModal) HandleDefaultSwitch(handleDefaultSwitch func(active int) (*tview.Application, int)) *MessageModal {
	modal.handleDefaultSwitch = handleDefaultSwitch
	return modal
}

func (modal *MessageModal) GetLayout(rowGap, columnGap int, spacer *tview.Box) (*tview.Flex, int) {
	var message *tview.TextView
	var success *tview.Button
	var failure *tview.Button

	var height int = 2
	var layout *tview.Flex
	var buttons *tview.Flex

	if modal.message == nil {
		modal.
			SetDefaultMessage("Message")
	}

	message = modal.message.
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
					modal.handleDefaultClick(MessageModalSuccessButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(MessageModalSuccessButton)

					if index == MessageModalFailureButton && failure != nil {
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
					modal.handleDefaultClick(MessageModalFailureButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(MessageModalFailureButton)

					if index == MessageModalSuccessButton && success != nil {
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
		AddItem(buttons, 1, 1, true)

	return layout, height
}
