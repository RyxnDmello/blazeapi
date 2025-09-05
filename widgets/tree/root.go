package tree

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Root struct {
	path                      string
	name                      string
	expanded                  bool
	padding                   int
	selectedIcon              rune
	unselectedIcon            rune
	selectedBackgroundColor   int32
	selectedForegroundColor   int32
	unselectedBackgroundColor int32
	unselectedForegroundColor int32
	handleSelect              func()
}

func NewRoot() *Root {
	return &Root{
		path:                      ".",
		name:                      ".",
		expanded:                  true,
		padding:                   0,
		selectedIcon:              '\uf07c',
		unselectedIcon:            '\uf07b',
		selectedBackgroundColor:   tcell.ColorDefault.Hex(),
		selectedForegroundColor:   tcell.ColorDefault.Hex(),
		unselectedBackgroundColor: tcell.ColorDefault.Hex(),
		unselectedForegroundColor: tcell.ColorDefault.Hex(),
		handleSelect:              nil,
	}
}

func (root *Root) SetPath(path string) *Root {
	root.path = path
	return root
}

func (root *Root) SetName(name string) *Root {
	root.name = name
	return root
}

func (root *Root) SetExpanded(expanded bool) *Root {
	root.expanded = expanded
	return root
}

func (root *Root) SetPadding(padding int) *Root {
	root.padding = max(0, padding)
	return root
}

func (root *Root) SetSelectedIcon(selectedIcon rune) *Root {
	root.selectedIcon = selectedIcon
	return root
}

func (root *Root) SetUnselectedIcon(unselectedIcon rune) *Root {
	root.unselectedIcon = unselectedIcon
	return root
}

func (root *Root) SetSelectedStyle(selectedBackgroundColor, selectedForegroundColor int32) *Root {
	root.selectedBackgroundColor = selectedBackgroundColor
	root.selectedForegroundColor = selectedForegroundColor
	return root
}

func (root *Root) SetUnselectedStyle(unselectedBackgroundColor, unselectedForegroundColor int32) *Root {
	root.unselectedBackgroundColor = unselectedBackgroundColor
	root.unselectedForegroundColor = unselectedForegroundColor
	return root
}

func (root *Root) HandleSelect(handleSelect func()) *Root {
	root.handleSelect = handleSelect
	return root
}

func (root *Root) GetPath() string {
	return root.path
}

func (root *Root) GetName() string {
	return root.name
}

func (root *Root) GetDisplay(selectedIcon, unselectedIcon bool) string {
	name := root.GetName()

	padding := root.padding - len(name)

	if selectedIcon {
		return string(root.GetSelectedIcon()) + " " + name + strings.Repeat(" ", padding-2)
	}

	if unselectedIcon {
		return string(root.GetUnselectedIcon()) + " " + name + strings.Repeat(" ", padding-2)
	}

	return name + strings.Repeat(" ", padding)
}

func (root *Root) GetSelectedIcon() rune {
	return root.selectedIcon
}

func (root *Root) GetUnselectedIcon() rune {
	return root.unselectedIcon
}

func (root *Root) GetSelectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(root.selectedBackgroundColor)).
		Foreground(tcell.NewHexColor(root.selectedForegroundColor))

	return style
}

func (root *Root) GetUnselectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(root.unselectedBackgroundColor)).
		Foreground(tcell.NewHexColor(root.unselectedForegroundColor))

	return style
}

func (root *Root) node() *Node {
	return &Node{
		parent:                    nil,
		path:                      root.path,
		name:                      root.name,
		dir:                       true,
		level:                     0,
		padding:                   root.padding,
		selectedIcon:              root.selectedIcon,
		unselectedIcon:            root.unselectedIcon,
		selectedBackgroundColor:   root.selectedBackgroundColor,
		selectedForegroundColor:   root.selectedForegroundColor,
		unselectedBackgroundColor: root.unselectedBackgroundColor,
		unselectedForegroundColor: root.unselectedForegroundColor,
		handleSelect:              root.handleSelect,
	}
}
