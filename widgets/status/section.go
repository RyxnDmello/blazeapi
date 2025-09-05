package status

import "github.com/rivo/tview"

type Section struct {
	badges    []*Badge
	separator *Glyph
}

func NewSection() *Section {
	return &Section{
		badges:    nil,
		separator: nil,
	}
}

func (section *Section) AddBadge(badge *Badge) *Section {
	section.badges = append(section.badges, badge)
	return section
}

func (section *Section) SetSeparator(separator Glyph) *Section {
	section.separator = &separator
	return section
}

func (section *Section) GetLayout() (*tview.Flex, int) {
	var width int = 0

	widget := tview.
		NewFlex().
		SetDirection(tview.FlexColumn)

	for index, badge := range section.badges {
		if section.separator != nil && index < len(section.badges)-1 {
			badge.
				AddRightGlyph(*section.separator)
		}

		width += badge.
			GetWidth()

		layout, width := badge.
			GetLayout()

		widget.
			AddItem(layout, width, 1, false)
	}

	return widget, width
}
