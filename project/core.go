package project

import (
	"os"
	"path/filepath"

	"blazeapi/core"
	"blazeapi/query"
	"blazeapi/response"
	"blazeapi/utils"
	"blazeapi/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeProject(app *tview.Application, query *query.Query, response *response.Response) (project *tview.TreeView, createFileModal *tview.Flex, createFolderModal *tview.Flex, deleteNodeModal *tview.Flex) {
	root := NewNode().
		Initialize(nil, "Test", "./test", true).
		Render()

	project = widgets.NewTree().
		SetRoot(root).
		SetTitle(" Manager").
		HandleSelect(
			func(treeNode *tview.TreeNode) {
				node, ok := treeNode.GetReference().(*Node)

				if !ok {
					return
				}

				if !node.Collection() {
					api := core.NewApi().Read(node.path)

					query.SetMethod(api.Method)
					query.SetUrl(api.Url)
					query.SetBody(api.Body)

					return
				}

				if treeNode.IsExpanded() {
					treeNode.Collapse().ClearChildren().SetExpanded(false)
					return
				}

				addDirectory(treeNode, node.path)

				treeNode.SetExpanded(true)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyCtrlA {
					collapseDirectory(project.GetRoot())
				}

				return event
			},
		).
		Render()

	createFolderModal = initializeCreateFolderModal(app, project)
	createFileModal = initializeCreateFileModal(app, project)
	deleteNodeModal = initializeDeleteNodeModal(app, project)

	addDirectory(root, "./test")

	return project, createFileModal, createFolderModal, deleteNodeModal
}

func initializeCreateFolderModal(app *tview.Application, project *tview.TreeView) (createFolderModal *tview.Flex) {
	var input *tview.InputField
	var button *tview.Button

	input = widgets.
		NewInput().
		SetPlaceholder("Enter Name").
		HandleAcceptance(
			func(text string, lastChar rune) bool {
				if !utils.ValidateIdentifier(text) {
					input.SetFieldTextColor(tcell.ColorRed)
					return false
				}

				input.SetFieldTextColor(tcell.ColorWhite)

				return true
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(button)
				}

				return event
			},
		).
		Render()

	button = widgets.NewButton().
		SetLabel("Create").
		HandleSelect(
			func() {
				if !utils.ValidateIdentifier(input.GetText()) {
					input.SetFieldTextColor(tcell.ColorRed)
					return
				}

				treeNode := project.GetCurrentNode()

				if treeNode == nil {
					return
				}

				node, ok := treeNode.GetReference().(*Node)

				if !ok {
					return
				}

				_, message, success := core.NewCollection().Create(input.GetText(), node.Path(true))

				input.SetText(message).SetFieldTextColor(tcell.ColorWhite)

				if !success {
					input.SetFieldTextColor(tcell.ColorRed)
				}

				if node.Collection() {
					expandDirectory(treeNode)
					return
				}

				expandDirectory(node.parent)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(input)
				}

				return event
			},
		).
		Render()

	createFolderModal = widgets.
		NewModal().
		SetTitle("Create Collection").
		SetDimension(50, 10).
		AddInput(input, true).
		AddButton(button, false).
		Render()

	return createFolderModal
}

func initializeCreateFileModal(app *tview.Application, project *tview.TreeView) (createFileModal *tview.Flex) {
	var input *tview.InputField
	var button *tview.Button

	input = widgets.
		NewInput().
		SetPlaceholder("Enter Name").
		HandleAcceptance(
			func(text string, lastChar rune) bool {
				if !utils.ValidateIdentifier(text) {
					input.SetFieldTextColor(tcell.ColorRed)
					return false
				}

				input.SetFieldTextColor(tcell.ColorWhite)

				return true
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(button)
				}

				return event
			},
		).
		Render()

	button = widgets.NewButton().
		SetLabel("Create").
		HandleSelect(
			func() {
				if !utils.ValidateIdentifier(input.GetText()) {
					input.SetFieldTextColor(tcell.ColorRed)
					return
				}

				treeNode := project.GetCurrentNode()

				if treeNode == nil {
					return
				}

				node, ok := treeNode.GetReference().(*Node)

				if !ok {
					return
				}

				_, message, success := core.NewApi().Create(input.GetText(), node.Path(true))

				input.SetText(message).SetFieldTextColor(tcell.ColorWhite)

				if !success {
					input.SetFieldTextColor(tcell.ColorRed)
				}

				if node.Collection() {
					expandDirectory(treeNode)
					return
				}

				expandDirectory(node.parent)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyTAB {
					app.SetFocus(input)
				}

				return event
			},
		).
		Render()

	createFileModal = widgets.
		NewModal().
		SetTitle("Create Request").
		SetDimension(50, 10).
		AddInput(input, true).
		AddButton(button, false).
		Render()

	return createFileModal
}

func initializeDeleteNodeModal(app *tview.Application, project *tview.TreeView) (deleteNodeModal *tview.Flex) {
	var message *tview.TextView
	var delete *tview.Button

	message = widgets.
		NewMessage().
		SetText("Are you sure want to delete?").
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				app.SetFocus(delete)
				return event
			},
		).
		Render()

	delete = widgets.
		NewButton().
		SetLabel("Delete").
		HandleSelect(
			func() {
				node, ok := project.GetCurrentNode().GetReference().(*Node)

				if !ok {
					return
				}

				if node.Collection() {
					core.NewCollection().Delete(node.path)
				}

				if !node.Collection() {
					core.NewApi().Delete(node.path)
				}

				expandDirectory(node.parent)
			},
		).
		HandleInput(
			func(event *tcell.EventKey) *tcell.EventKey {
				return event
			},
		).
		Render()

	deleteNodeModal = widgets.
		NewModal().
		SetTitle("Add Artifact").
		AddInput(message, true).
		AddButton(delete, true).
		SetDimension(50, 10).
		Render()

	return deleteNodeModal
}

func addDirectory(parent *tview.TreeNode, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Invalid Collection")
	}

	for _, entry := range entries {
		name := entry.Name()

		node := NewNode().
			Initialize(parent, name, filepath.Join(path, name), entry.IsDir()).
			Render()

		parent.AddChild(node)
	}
}

func expandDirectory(treeNode *tview.TreeNode) {
	if treeNode == nil {
		return
	}

	node, ok := treeNode.GetReference().(*Node)

	if !ok {
		return
	}

	if node.Collection() {
		treeNode.Collapse().ClearChildren().Expand()
		addDirectory(treeNode, node.path)
		return
	}

	node.parent.Collapse().ClearChildren().Expand()

	addDirectory(node.parent, node.ParentNode().path)
}

func collapseDirectory(treeNode *tview.TreeNode) {
	if treeNode == nil {
		return
	}

	node, ok := treeNode.GetReference().(*Node)

	if !ok {
		return
	}

	if node.Collection() {
		treeNode.Collapse().ClearChildren()
		return
	}

	node.parent.Collapse().ClearChildren()
}
