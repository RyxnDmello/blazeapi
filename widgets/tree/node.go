package tree

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Node struct {
	parent                    *Node
	path                      string
	name                      string
	dir                       bool
	level                     int
	padding                   int
	selectedIcon              rune
	unselectedIcon            rune
	selectedBackgroundColor   int32
	selectedForegroundColor   int32
	unselectedBackgroundColor int32
	unselectedForegroundColor int32
	handleSelect              func()
}

func NewNode() *Node {
	return &Node{
		parent:                    nil,
		path:                      ".",
		name:                      ".",
		dir:                       true,
		level:                     0,
		padding:                   0,
		selectedIcon:              rune(0),
		unselectedIcon:            rune(0),
		selectedBackgroundColor:   tcell.ColorDefault.Hex(),
		selectedForegroundColor:   tcell.ColorDefault.Hex(),
		unselectedBackgroundColor: tcell.ColorDefault.Hex(),
		unselectedForegroundColor: tcell.ColorDefault.Hex(),
		handleSelect:              nil,
	}
}

func (node *Node) SetParent(parent *Node) *Node {
	node.parent = parent
	return node
}

func (node *Node) SetPath(path string) *Node {
	node.path = path
	return node
}

func (node *Node) SetName(name string) *Node {
	node.name = name
	return node
}

func (node *Node) SetDir(dir bool) *Node {
	node.dir = dir
	return node
}

func (node *Node) SetLevel(level int) *Node {
	node.level = max(0, level)
	return node
}

func (node *Node) SetPadding(padding int) *Node {
	node.padding = max(0, padding)
	return node
}

func (node *Node) SetSelectedIcon(selectedIcon rune) *Node {
	node.selectedIcon = selectedIcon
	return node
}

func (node *Node) SetUnselectedIcon(unselectedIcon rune) *Node {
	node.unselectedIcon = unselectedIcon
	return node
}

func (node *Node) SetSelectedStyle(selectedBackgroundColor, selectedForegroundColor int32) *Node {
	node.selectedBackgroundColor = selectedBackgroundColor
	node.selectedForegroundColor = selectedForegroundColor
	return node
}

func (node *Node) SetUnselectedStyle(unselectedBackgroundColor, unselectedForegroundColor int32) *Node {
	node.unselectedBackgroundColor = unselectedBackgroundColor
	node.unselectedForegroundColor = unselectedForegroundColor
	return node
}

func (node *Node) HandleSelect(handleSelect func()) *Node {
	node.handleSelect = handleSelect
	return node
}

func (node *Node) GetParent() *Node {
	return node.parent
}

func (node *Node) GetPath() string {
	return node.path
}

func (node *Node) GetName() string {
	return strings.Split(node.name, ".")[0]
}

func (node *Node) GetExtension() string {
	splits := strings.Split(node.name, ".")

	if len(splits) == 2 {
		return "." + splits[1]
	}

	return ""
}

func (node *Node) GetDisplay(selectedIcon, unselectedIcon, extension bool) string {
	name := node.GetName()

	if extension {
		name += node.GetExtension()
	}

	padding := node.padding - len(name) - (2 * node.level)

	if selectedIcon {
		return string(node.GetSelectedIcon()) + " " + name + strings.Repeat(" ", padding-2)
	}

	if unselectedIcon {
		return string(node.GetUnselectedIcon()) + " " + name + strings.Repeat(" ", padding-2)
	}

	return name + strings.Repeat(" ", padding)
}

func (node *Node) GetLevel() int {
	return node.level
}

func (node *Node) GetPadding() int {
	return node.padding
}

func (node *Node) GetSelectedIcon() rune {
	return node.selectedIcon
}

func (node *Node) GetUnselectedIcon() rune {
	return node.unselectedIcon
}

func (node *Node) GetSelectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(node.selectedBackgroundColor)).
		Foreground(tcell.NewHexColor(node.selectedForegroundColor))

	return style
}

func (node *Node) GetUnselectedStyle() tcell.Style {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(node.unselectedBackgroundColor)).
		Foreground(tcell.NewHexColor(node.unselectedForegroundColor))

	return style
}

func (node *Node) GetSelectedStyleDecompose() (int32, int32) {
	selectedForegroundColor, selectedBackgroundColor, _ := node.GetSelectedStyle().Decompose()
	return selectedBackgroundColor.Hex(), selectedForegroundColor.Hex()
}

func (node *Node) GetUnselectedStyleDecompose() (int32, int32) {
	unselectedForegroundColor, unselectedBackgroundColor, _ := node.GetUnselectedStyle().Decompose()
	return unselectedBackgroundColor.Hex(), unselectedForegroundColor.Hex()
}

func (node *Node) IsDir() bool {
	return node.dir
}

func (node *Node) Draw() *tview.TreeNode {
	display := node.
		GetDisplay(false, true, true)

	widget := tview.
		NewTreeNode(display).
		SetText(display).
		SetReference(node).
		SetExpanded(false).
		SetSelectable(true).
		SetTextStyle(node.GetUnselectedStyle()).
		SetSelectedTextStyle(node.GetSelectedStyle())

	widget.
		SetSelectedFunc(func() {
			if node.handleSelect != nil {
				node.handleSelect()
			}
		})

	return widget
}
