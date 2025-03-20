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

// NewTree creates and returns an instance of the tree type.
//
// The tree widget is used to display hierarchical data in Blaze. Properties
// such as the title, root node, and event handlers for selection and input
// can be customized. The widget aligns with the features and requirements
// of Blaze. Under the hood, it uses tview.TreeView to render the user interface.
//
// # Properties
//
//	title        string
//	root         *tview.TreeNode
//	handleSelect func(node *tview.TreeNode)
//	handleInput  func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*tree
//
// # Usage
//
//	tree := widgets.NewTree()
func NewTree() *tree {
	return &tree{
		title:        "",
		root:         nil,
		handleSelect: nil,
		handleInput:  nil,
	}
}

// SetTitle sets the title of the tree type.
//
// The title is displayed at the top of the tree view. Method chaining is
// supported.
//
// # Parameters
//
//	title string
//
// # Returns
//
//	*tree
//
// # Usage
//
//	tree := widgets.NewTree().SetTitle("")
func (project *tree) SetTitle(title string) *tree {
	project.title = title
	return project
}

// SetRoot sets the root node of the tree type.
//
// The root node defines the starting point of the tree hierarchy. Method
// chaining is supported.
//
// # Parameters
//
//	root *tview.TreeNode
//
// # Returns
//
//	*tree
//
// # Usage
//
//	tree := widgets.NewTree().SetRoot(nil)
func (project *tree) SetRoot(root *tview.TreeNode) *tree {
	project.root = root
	return project
}

// HandleSelect assigns a custom function to handle node selection.
//
// Selection behavior is customized with a function that runs when a node is
// selected. Method chaining is supported.
//
// # Parameters
//
//	handleSelect func(node *tview.TreeNode)
//
// # Returns
//
//	*tree
//
// # Usage
//
//	tree := widgets.
//	     NewTree().
//	     HandleSelect(
//	         func(node *tview.TreeNode) {
//	         },
//	     )
func (project *tree) HandleSelect(handleSelect func(node *tview.TreeNode)) *tree {
	project.handleSelect = handleSelect
	return project
}

// HandleInput assigns a custom function to process keyboard events.
//
// Keyboard behavior can be customized with a tcell EventKey handler. If nil,
// tview defaults are used. Method chaining is supported.
//
// # Parameters
//
//	handleInput func(event *tcell.EventKey) *tcell.EventKey
//
// # Returns
//
//	*tree
//
// # Usage
//
//	tree := widgets.
//	     NewTree().
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     )
func (project *tree) HandleInput(handleInput func(event *tcell.EventKey) *tcell.EventKey) *tree {
	project.handleInput = handleInput
	return project
}

// Render creates and returns a tview.TreeView with the configured properties.
//
// A tview.TreeView is built with the set properties. This function renders
// the tree view for display in Blaze.
//
// # Returns
//
//	*tview.TreeView
//
// # Usage
//
//	tree := widgets.
//	     NewTree().
//	     SetTitle("").
//	     SetRoot(nil).
//	     HandleSelect(
//	         func(node *tview.TreeNode) {
//	         },
//	     ).
//	     HandleInput(
//	         func(event *tcell.EventKey) *tcell.EventKey {
//	             return event
//	         },
//	     ).
//	     Render()
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

// title formats the title string for display.
//
// A formatted string is returned with the title enclosed in spaces. This
// function is a helper for styling tree titles.
//
// # Parameters
//
//	title string
//
// # Returns
//
//	string
//
// # Usage
//
//	title := widgets.title("")
func title(title string) string {
	return fmt.Sprintf(" %s ", title)
}
