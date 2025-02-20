package project

import (
	"os"
	"path/filepath"
	"strings"

	"blazeapi/core"
	"blazeapi/query"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeProject(app *tview.Application, query query.Query, response response.Response) (project *tview.TreeView, createNodeModal *tview.Flex, deleteNodeModal *tview.Flex) {
	node := CreateNode(nil, "Test", "./test", true)

	root := tview.
		NewTreeNode(node.Name(true, false)).
		SetReference(node)

	createDirectory(root, "./test")

	project = tview.
		NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphicsColor(tcell.NewRGBColor(75, 75, 75))

	createNodeModal = initializeCreateNodeModal(app, project)
	deleteNodeModal = initializeDeleteNodeModal(project)

	project.SetSelectedFunc(func(treeNode *tview.TreeNode) {
		node, ok := treeNode.GetReference().(Node)

		if !ok {
			return
		}

		if !node.IsCollection() {
			api := CreateAPI(node.path)

			query.SetMethod(api.Method)
			query.SetUrl(api.Url)

			return
		}

		if treeNode.IsExpanded() {
			treeNode.SetExpanded(false).ClearChildren()
			return
		}

		createDirectory(treeNode, node.path)

		treeNode.SetExpanded(true)
	})

	project.
		SetBorder(true).
		SetTitle("  Manager ").
		SetBorderPadding(1, 0, 1, 0).
		SetTitleAlign(tview.AlignLeft)

	return project, createNodeModal, deleteNodeModal
}

func initializeCreateNodeModal(app *tview.Application, project *tview.TreeView) (createNodeModal *tview.Flex) {
	var input *tview.InputField
	var createAPI *tview.Button
	var createFolder *tview.Button

	input = Input(
		"Enter Name",
		func(textToCheck string, lastChar rune) bool {
			if !ValidateName(textToCheck) {
				input.SetFieldTextColor(tcell.ColorRed)
				return false
			}

			input.SetFieldTextColor(tcell.ColorWhite)

			return true
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(createAPI)
			}

			return event
		},
	)

	createAPI = Button(
		"Add Request",
		func() {
			if strings.HasSuffix(input.GetText(), "_") {
				input.SetFieldTextColor(tcell.ColorRed)
				return
			}

			treeNode := project.GetCurrentNode()

			if treeNode == nil {
				return
			}

			node, ok := treeNode.GetReference().(Node)

			if !ok {
				return
			}

			message, success := core.AddAPI(input.GetText(), node.Path(true))

			if success {
				input.SetText(message)
			}

			if node.IsCollection() {
				if !treeNode.IsExpanded() {
					collapseDirectory(treeNode)
					return
				}

				expandDirectory(treeNode)
				return
			}

			expandDirectory(node.parent)
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(createFolder)
			}

			return event
		},
	)

	createFolder = Button(
		"Add Folder",
		func() {
			if strings.HasSuffix(input.GetText(), "_") {
				input.SetFieldTextColor(tcell.ColorRed)
				return
			}

			treeNode := project.GetCurrentNode()

			if treeNode == nil {
				return
			}

			node, ok := treeNode.GetReference().(Node)

			if !ok {
				return
			}

			message, success := core.AddCollection(input.GetText(), node.Path(true))

			if success {
				input.SetText(message)
			}

			if node.IsCollection() {
				if !treeNode.IsExpanded() {
					collapseDirectory(treeNode)
					return
				}

				expandDirectory(treeNode)
				return
			}

			expandDirectory(node.parent)
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTAB {
				app.SetFocus(input)
			}

			return event
		},
	)

	layout := tview.
		NewGrid().
		AddItem(input, 0, 0, 1, 2, 0, 0, true).
		AddItem(createAPI, 1, 0, 1, 1, 0, 0, false).
		AddItem(createFolder, 1, 1, 1, 1, 0, 0, false)

	layout.
		SetBorder(true).
		SetTitle("  Create Artifact ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 1)

	alignment := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(layout, 8, 1, true).
		AddItem(nil, 0, 1, false)

	createNodeModal = tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 50, 1, true).
		AddItem(nil, 0, 1, false)

	return createNodeModal
}

func initializeDeleteNodeModal(project *tview.TreeView) (createNodeModal *tview.Flex) {
	var text *tview.TextView
	var delete *tview.Button

	text = Display("Are you sure want to delete?")

	delete = Button(
		"Delete",
		func() {
			node, ok := project.GetCurrentNode().GetReference().(Node)

			if !ok {
				return
			}

			if node.IsCollection() {
				core.DeleteCollection(node.path)
			}

			if !node.IsCollection() {
				core.DeleteAPI(node.path)
			}

			expandDirectory(node.parent)
		},
		func(event *tcell.EventKey) *tcell.EventKey {
			return event
		},
	)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, false).
		AddItem(delete, 0, 1, true)

	layout.
		SetBorder(true).
		SetTitle("  Create Artifact ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 1, 1)

	alignment := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(layout, 8, 1, true).
		AddItem(nil, 0, 1, false)

	createNodeModal = tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 50, 1, true).
		AddItem(nil, 0, 1, false)

	return createNodeModal
}

func createDirectory(parent *tview.TreeNode, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Invalid Collection")
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(path, name)

		node := CreateNode(parent, name, path, entry.IsDir())

		child := tview.
			NewTreeNode(node.Name(true, false)).
			SetExpanded(false).
			SetReference(node)

		parent.AddChild(child)
	}
}

func expandDirectory(treeNode *tview.TreeNode) {
	if treeNode == nil {
		return
	}

	node, ok := treeNode.GetReference().(Node)

	if !ok {
		return
	}

	if node.IsCollection() {
		treeNode.Collapse().ClearChildren().Expand()
		createDirectory(treeNode, node.path)
		return
	}

	node.parent.Collapse().ClearChildren().Expand()
	createDirectory(node.parent, node.ParentNode().path)
}

func collapseDirectory(treeNode *tview.TreeNode) {
	if treeNode == nil {
		return
	}

	node, ok := treeNode.GetReference().(Node)

	if !ok {
		return
	}

	if node.IsCollection() {
		treeNode.Collapse().ClearChildren()
		return
	}

	node.parent.Collapse().ClearChildren()
}
