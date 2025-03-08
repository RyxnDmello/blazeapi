package project

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Project struct {
	title        string
	root         *tview.TreeNode
	handleSelect func(node *tview.TreeNode)
	handleInput  func(event *tcell.EventKey) *tcell.EventKey
}

func NewProject() *Project {
	return &Project{
		title:        "",
		root:         nil,
		handleSelect: nil,
		handleInput:  nil,
	}
}

func (project *Project) SetTitle(title string) *Project {
	project.title = title
	return project
}

func (project *Project) SetRoot(root *tview.TreeNode) *Project {
	project.root = root
	return project
}

func (project *Project) HandleSelect(handleSelect func(node *tview.TreeNode)) *Project {
	project.handleSelect = handleSelect
	return project
}

func (project *Project) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *Project {
	project.handleInput = handleInput
	return project
}

func (project *Project) Render() *tview.TreeView {
	tree := tview.
		NewTreeView().
		SetRoot(project.root).
		SetCurrentNode(project.root).
		SetSelectedFunc(project.handleSelect).
		SetGraphicsColor(tcell.NewRGBColor(75, 75, 75))

	tree.
		SetBorder(true).
		SetTitle(project.title).
		SetBorderPadding(1, 0, 1, 0).
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(project.handleInput)

	return tree
}
