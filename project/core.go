package project

import (
	"os"
	"path/filepath"
	"strings"

	"blazeapi/core"
	"blazeapi/query"
	"blazeapi/response"
	"blazeapi/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeProject(app *tview.Application, query query.Query, response response.Response) (project *tview.TreeView, createNodeModal *tview.Flex, deleteNodeModal *tview.Flex) {
	root := NewNode().
		Initialize(nil, "Test", "./test", true).
		Render()

	project = NewProject().
		SetRoot(root).
		SetTitle("  Manager ").
		HandleSelect(
			func(treeNode *tview.TreeNode) {
				node, ok := treeNode.GetReference().(*Node)

				if !ok {
					return
				}

				if !node.Collection() {
					api := CreateAPI(node.path)

					query.SetMethod(api.Method)
					query.SetUrl(api.Url)

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

	createNodeModal = initializeCreateNodeModal(app, project)
	deleteNodeModal = initializeDeleteNodeModal(app, project)

	addDirectory(root, "./test")

	return project, createNodeModal, deleteNodeModal
}

func initializeCreateNodeModal(app *tview.Application, project *tview.TreeView) (createNodeModal *tview.Flex) {
	var input *tview.InputField
	var request *tview.Button
	var folder *tview.Button

	input = widgets.
		NewInput().
		SetPlaceholder("Enter Name").
		HandleAcceptance(
			func(text string, lastChar rune) bool {
				if !ValidateName(text) {
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
					app.SetFocus(request)
				}

				return event
			},
		).
		Render()

	request = widgets.
		NewButton().
		SetLabel("Add Folder").
		HandleSelect(
			func() {
				if strings.HasSuffix(input.GetText(), "_") {
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

				message, success := core.AddAPI(input.GetText(), node.Path(true))

				if success {
					input.SetText(message)
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
					app.SetFocus(folder)
				}

				return event
			},
		).
		Render()

	folder = widgets.NewButton().
		SetLabel("Add Folder").
		HandleSelect(
			func() {
				if strings.HasSuffix(input.GetText(), "_") {
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

				message, success := core.AddCollection(input.GetText(), node.Path(true))

				if success {
					input.SetText(message)
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

	createNodeModal = widgets.
		NewModal().
		SetTitle("Add Artifact").
		SetDimension(50, 10).
		AddInput(input, true).
		AddButton(request, false).
		AddButton(folder, false).
		Render()

	return createNodeModal
}

func initializeDeleteNodeModal(app *tview.Application, project *tview.TreeView) (deleteNodeModal *tview.Flex) {
	var message *tview.TextView
	var delete *tview.Button

	message = widgets.
		NewMessage().
		SetLabel("Are you sure want to delete?").
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
				node, ok := project.GetCurrentNode().GetReference().(Node)

				if !ok {
					return
				}

				if node.Collection() {
					core.DeleteCollection(node.path)
				}

				if !node.Collection() {
					core.DeleteAPI(node.path)
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
		AddInput(message, false).
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
