package tree

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tree struct {
	root            *Root
	topmost         bool
	hierarchy       bool
	design          Design
	handleChange    func(node *Node, treeNode *tview.TreeNode)
	handleSelect    func(node *Node, treeNode *tview.TreeNode)
	handleSwitch    func()
	handleInput     func(event *tcell.EventKey) *tcell.EventKey
	handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)
}

func NewTree() *Tree {
	return &Tree{
		root:            nil,
		topmost:         true,
		hierarchy:       true,
		design:          NewDesign(),
		handleChange:    nil,
		handleSelect:    nil,
		handleSwitch:    nil,
		handleInput:     nil,
		handleDimension: nil,
	}
}

func (tree *Tree) SetRoot(root *Root) *Tree {
	tree.root = root
	return tree
}

func (tree *Tree) SetTopmost(topmost bool) *Tree {
	tree.topmost = topmost
	return tree
}

func (tree *Tree) SetHierarchy(hierarchy bool) *Tree {
	tree.hierarchy = hierarchy
	return tree
}

func (tree *Tree) SetDesign(design Design) *Tree {
	tree.design = design
	return tree
}

func (tree *Tree) HandleChange(handleChange func(node *Node, treeNode *tview.TreeNode)) *Tree {
	tree.handleChange = handleChange
	return tree
}

func (tree *Tree) HandleSelect(handleSelect func(node *Node, treeNode *tview.TreeNode)) *Tree {
	tree.handleSelect = handleSelect
	return tree
}

func (tree *Tree) HandleSwitch(handleSwitch func()) *Tree {
	tree.handleSwitch = handleSwitch
	return tree
}

func (tree *Tree) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Tree {
	tree.handleInput = handleInput
	return tree
}

func (tree *Tree) HandleDimension(handleDimension func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Tree {
	tree.handleDimension = handleDimension
	return tree
}

func (tree *Tree) GetRoot() *Root {
	return tree.root
}

func (tree *Tree) GetRootNode() *Node {
	return tree.root.node()
}

func (tree *Tree) GetRootView() *tview.TreeNode {
	return tree.GetRootNode().Draw().SetExpanded(true).SetSelectable(true)
}

func (tree *Tree) GetTopmost() int {
	if tree.topmost {
		return 0
	}

	return 1
}

func (tree *Tree) View() *tview.TreeView {
	root := tree.
		GetRootView()

	widget := tview.
		NewTreeView().
		SetRoot(root).
		SetGraphics(false).
		SetCurrentNode(root).
		SetAlign(!tree.hierarchy).
		SetTopLevel(tree.GetTopmost())

	widget.
		SetChangedFunc(func(treeNode *tview.TreeNode) {
			if tree.handleChange == nil {
				return
			}

			node, ok := treeNode.GetReference().(*Node)

			if !ok {
				return
			}

			tree.handleChange(node, treeNode)
		}).
		SetSelectedFunc(func(treeNode *tview.TreeNode) {
			if tree.handleSelect == nil {
				return
			}

			node, ok := treeNode.GetReference().(*Node)

			if !ok {
				return
			}

			tree.handleSelect(node, treeNode)
		}).
		SetDoneFunc(func(key tcell.Key) {
			if tree.handleSwitch != nil && key == tcell.KeyTab {
				tree.handleSwitch()
				return
			}
		})

	widget.
		SetBorder(false).
		SetBorderPadding(0, 0, 0, 0).
		SetBackgroundColor(tree.design.GetRootColor())

	widget.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if tree.handleSwitch != nil && event.Key() == tcell.KeyTab {
				tree.handleSwitch()
			}

			if tree.handleInput != nil {
				return tree.handleInput(event)
			}

			return event
		}).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			for _, glyph := range tree.design.GetGlyphs() {
				glyph.Render(widget.Box, screen, widget.HasFocus(), false)
			}

			if tree.handleDimension == nil {
				return x, y, width, height
			}

			return tree.handleDimension(screen, x, y, width, height)
		})

	AddDirectory(root)

	return widget
}
