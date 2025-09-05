package modal

import (
	"blazeapi/widgets/area"
	"blazeapi/widgets/button"
	"blazeapi/widgets/glyph"
	"blazeapi/widgets/message"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	AreaModalArea          = 0
	AreaModalSuccessButton = 1
	AreaModalFailureButton = 2
)

type AreaModal struct {
	message             *message.Message
	area                *area.Area
	success             *button.Button
	failure             *button.Button
	handleDefaultClick  func(active int)
	handleDefaultSwitch func(active int) (*tview.Application, int)
	handleDefaultChange func(value string, length int, firstChar, lastChar rune)
}

func NewAreaModal() *AreaModal {
	return &AreaModal{
		message:             nil,
		area:                nil,
		success:             nil,
		failure:             nil,
		handleDefaultClick:  nil,
		handleDefaultSwitch: nil,
		handleDefaultChange: nil,
	}
}

func (modal *AreaModal) SetMessage(message *message.Message) *AreaModal {
	modal.message = message
	return modal
}

func (modal *AreaModal) SetDefaultMessage(text string) *AreaModal {
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

func (modal *AreaModal) SetArea(area *area.Area) *AreaModal {
	modal.area = area
	return modal
}

func (modal *AreaModal) SetDefaultArea(label, placeholder string, rows, columns int) *AreaModal {
	design := area.
		NewDesign().
		SetRootColor(0x0e0e15).
		SetTextStyle(0x0e0e15, 0xcdd6f4).
		SetLabelStyle(0x0e0e15, 0xcdd6f4).
		SetPlaceholderStyle(0x0e0e15, 0xcdd6f4)

	modal.area = area.
		NewTextArea().
		SetLabel(label).
		SetDesign(design).
		SetSize(rows, columns).
		SetPlaceholder(placeholder)

	return modal
}

func (modal *AreaModal) SetSuccessButton(success *button.Button) *AreaModal {
	modal.success = success
	return modal
}

func (modal *AreaModal) SetDefaultSuccessButton(label string) *AreaModal {
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

func (modal *AreaModal) SetFailureButton(failure *button.Button) *AreaModal {
	modal.failure = failure
	return modal
}

func (modal *AreaModal) SetDefaultFailureButton(label string) *AreaModal {
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

func (modal *AreaModal) HandleDefaultClick(handleDefaultClick func(active int)) *AreaModal {
	modal.handleDefaultClick = handleDefaultClick
	return modal
}

func (modal *AreaModal) HandleDefaultSwitch(handleDefaultSwitch func(active int) (*tview.Application, int)) *AreaModal {
	modal.handleDefaultSwitch = handleDefaultSwitch
	return modal
}

func (modal *AreaModal) HandleDefaultChange(handleDefaultChange func(value string, length int, firstChar, lastChar rune)) *AreaModal {
	modal.handleDefaultChange = handleDefaultChange
	return modal
}

func (modal *AreaModal) GetLayout(rowGap, columnGap int, spacer *tview.Box) (*tview.Flex, int) {
	var message *tview.TextView
	var success *tview.Button
	var failure *tview.Button
	var area *tview.TextArea

	var height int = 3
	var layout *tview.Flex
	var buttons *tview.Flex

	if modal.message == nil {
		modal.
			SetDefaultMessage("Message")
	}

	message = modal.message.
		View()

	if modal.area == nil {
		modal.
			SetDefaultArea("", "Enter Message", 0, 0)
	}

	handleChange := modal.area.
		GetHandleChange()

	handleSwitch := modal.area.
		GetHandleSwitch()

	area = modal.area.
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
				app, active := modal.handleDefaultSwitch(AreaModalArea)

				if active == AreaModalSuccessButton && success != nil {
					app.SetFocus(success)
				}

				if active == AreaModalFailureButton && failure != nil {
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
					modal.handleDefaultClick(AreaModalSuccessButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(AreaModalSuccessButton)

					if index == AreaModalArea {
						app.SetFocus(area)
					}

					if index == AreaModalFailureButton && failure != nil {
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
					modal.handleDefaultClick(AreaModalFailureButton)
					return
				}

				if handleClick != nil {
					handleClick()
				}
			}).
			HandleSwitch(func() {
				if modal.handleDefaultSwitch != nil {
					app, index := modal.handleDefaultSwitch(AreaModalFailureButton)

					if index == AreaModalArea {
						app.SetFocus(area)
					}

					if index == AreaModalSuccessButton && success != nil {
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
		AddItem(area, modal.area.GetHeight(), 1, true)

	height += modal.area.GetHeight() - 1

	if rowGap > 0 {
		layout.
			AddItem(spacer, rowGap, 1, false)

		height += rowGap
	}

	layout.
		AddItem(buttons, 1, 1, true)

	return layout, height
}
