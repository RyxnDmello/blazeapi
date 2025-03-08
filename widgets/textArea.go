package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type textArea struct {
	label       string
	placeholder string
	rows        int
	columns     int
	wordWrap    bool
	handleInput func(event *tcell.EventKey) *tcell.EventKey
}

func NewTextArea() *textArea {
	return &textArea{
		label:       "",
		placeholder: "",
		rows:        0,
		columns:     0,
		wordWrap:    true,
		handleInput: nil,
	}
}

func (textArea *textArea) SetLabel(label string) *textArea {
	textArea.label = label
	return textArea
}

func (textArea *textArea) SetPlaceholder(placeholder string) *textArea {
	textArea.placeholder = placeholder
	return textArea
}

func (textArea *textArea) SetRows(rows int) *textArea {
	textArea.rows = rows
	return textArea
}

func (textArea *textArea) SetColumns(columns int) *textArea {
	textArea.columns = columns
	return textArea
}

func (textArea *textArea) SetWordWrap(wordWrap bool) *textArea {
	textArea.wordWrap = wordWrap
	return textArea
}

func (textArea *textArea) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *textArea {
	textArea.handleInput = handleInput
	return textArea
}

func (textArea *textArea) Render() *tview.TextArea {
	area := tview.
		NewTextArea().
		SetLabel(textArea.label).
		SetWordWrap(textArea.wordWrap).
		SetPlaceholder(textArea.placeholder).
		SetSize(textArea.rows, textArea.columns).
		SetTextStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorWhite),
		).
		SetLabelStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		).
		SetPlaceholderStyle(
			tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.NewRGBColor(200, 200, 200)),
		)

	area.
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1).
		SetBackgroundColor(tcell.ColorBlack).
		SetInputCapture(textArea.handleInput)

	return area
}
