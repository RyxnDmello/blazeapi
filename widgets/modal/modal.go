package modal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ModalContent interface {
	GetLayout(rowGap, columnGap int, spacer *tview.Box) (*tview.Flex, int)
}

type Modal[M ModalContent] struct {
	content         M
	width           int
	height          int
	rowGap          int
	columnGap       int
	topPadding      int
	leftPadding     int
	rightPadding    int
	bottomPadding   int
	backgroundColor int32
}

func NewModal[M ModalContent](content M) *Modal[M] {
	return &Modal[M]{
		content:         content,
		width:           50,
		height:          0,
		rowGap:          1,
		columnGap:       1,
		topPadding:      1,
		leftPadding:     2,
		rightPadding:    2,
		bottomPadding:   1,
		backgroundColor: tcell.ColorDefault.Hex(),
	}
}

func (modal *Modal[M]) SetWidth(width int) *Modal[M] {
	modal.width = max(40, width)
	return modal
}

func (modal *Modal[M]) SetHeight(height int) *Modal[M] {
	modal.height = max(0, height)
	return modal
}

func (modal *Modal[M]) SetGap(rowGap, columnGap int) *Modal[M] {
	modal.rowGap = max(0, rowGap)
	modal.columnGap = max(0, columnGap)
	return modal
}

func (modal *Modal[M]) SetPadding(topPadding, leftPadding, rightPadding, bottomPadding int) *Modal[M] {
	modal.topPadding = max(0, topPadding)
	modal.leftPadding = max(0, leftPadding)
	modal.rightPadding = max(0, rightPadding)
	modal.bottomPadding = max(0, bottomPadding)
	return modal
}

func (modal *Modal[M]) SetVerticalPadding(verticalPadding int) *Modal[M] {
	modal.topPadding = max(0, verticalPadding)
	modal.bottomPadding = max(0, verticalPadding)
	return modal
}

func (modal *Modal[M]) SetHorizontalPadding(horizontalPadding int) *Modal[M] {
	modal.leftPadding = max(0, horizontalPadding)
	modal.rightPadding = max(0, horizontalPadding)
	return modal
}

func (modal *Modal[M]) SetBackgroundColor(backgroundColor int32) *Modal[M] {
	modal.backgroundColor = backgroundColor
	return modal
}

func (modal *Modal[M]) GetBackgroundColor() tcell.Color {
	return tcell.NewHexColor(modal.backgroundColor)
}

func (modal *Modal[M]) GetSpacer() *tview.Box {
	box := tview.
		NewBox().
		SetBackgroundColor(modal.GetBackgroundColor())

	return box
}

func (modal *Modal[M]) GetPopup(layout *tview.Flex, height int) *tview.Flex {
	if modal.height != 0 {
		height = modal.height
	}

	spacer := modal.
		GetSpacer()

	innerHorizontal := tview.
		NewFlex().
		SetDirection(tview.FlexColumn)

	if modal.leftPadding > 0 {
		innerHorizontal.
			AddItem(spacer, modal.leftPadding, 1, false)
	}

	innerHorizontal.
		AddItem(layout, 0, 1, true)

	if modal.rightPadding > 0 {
		innerHorizontal.
			AddItem(spacer, modal.rightPadding, 1, false)
	}

	vertical := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false)

	if modal.topPadding > 0 {
		vertical.
			AddItem(spacer, modal.topPadding, 1, false)
	}

	vertical.
		AddItem(innerHorizontal, height, 1, true)

	if modal.bottomPadding > 0 {
		vertical.
			AddItem(spacer, modal.bottomPadding, 1, false)
	}

	vertical.
		AddItem(nil, 0, 1, false)

	horizontal := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(vertical, modal.width, 1, true).
		AddItem(nil, 0, 1, false)

	return horizontal
}

func (modal *Modal[M]) View() *tview.Flex {
	spacer := modal.
		GetSpacer()

	layout, height := modal.content.
		GetLayout(modal.rowGap, modal.columnGap, spacer)

	popup := modal.
		GetPopup(layout, height)

	return popup
}
