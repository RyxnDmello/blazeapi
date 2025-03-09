package widgets

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type tree struct {
	title        string
	root         *tview.TreeNode
	handleSelect func(node *tview.TreeNode)
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

func NewTree() *tree {
	return &tree{
		title:        "",
		root:         nil,
		handleSelect: nil,
		handleInput:  nil,
	}
}

func (project *tree) SetTitle(title string) *tree {
	project.title = title
	return project
}

func (project *tree) SetRoot(root *tview.TreeNode) *tree {
	project.root = root
	return project
}

func (project *tree) HandleSelect(handleSelect func(node *tview.TreeNode)) *tree {
	project.handleSelect = handleSelect
	return project
}

func (project *tree) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *tree {
	project.handleInput = handleInput
	return project
}

func (project *tree) Render() *tview.TreeView {
	tree := tview.
		NewTreeView().
		SetRoot(project.root).
		SetCurrentNode(project.root).
		SetSelectedFunc(project.handleSelect).
		SetGraphicsColor(tcell.NewRGBColor(75, 75, 75))

	tree.
		SetBorder(true).
		SetBorderPadding(1, 0, 1, 0).
		SetTitle(title(project.title)).
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(project.handleInput)

	return tree
}

func title(title string) string {
	return fmt.Sprintf(" %s ", title)
}
