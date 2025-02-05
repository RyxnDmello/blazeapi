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

func InitializeProject(app *tview.Application, query query.Query, response response.Response) (project *tview.TreeView, projectCreateModal *tview.Flex) {
	root := tview.NewTreeNode("./test")

	createDirectory(root, "./test")

	project = tview.
		NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphics(false).
		SetTopLevel(1)

	projectCreateModal = initializeProjectCreateModal(app, project)

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

	return project, projectCreateModal
}

func initializeProjectCreateModal(app *tview.Application, project *tview.TreeView) (projectCreateModal *tview.Flex) {
	var input *tview.InputField
	var createAPI *tview.Button
	var createFolder *tview.Button

	input = Input(
		"Enter Name",
		func(textToCheck string, lastChar rune) bool {
			if !ValidateFileIdentifier(textToCheck) {
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
			treeNode := project.GetCurrentNode()

			node, ok := treeNode.GetReference().(Node)

			if !ok {
				return
			}

			name := input.GetText()

			if strings.HasSuffix(name, "_") {
				input.SetFieldTextColor(tcell.ColorRed)
				return
			}

			_, success := core.AddAPI(name, node.Path(true))

			if success {
				input.SetText("")
			}
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
			treeNode := project.GetCurrentNode()

			node, ok := treeNode.GetReference().(Node)

			if !ok {
				return
			}

			name := input.GetText()

			if strings.HasSuffix(name, "_") {
				input.SetFieldTextColor(tcell.ColorRed)
				return
			}

			_, success := core.AddCollection(name, node.Path(true))

			if success {
				input.SetText("")
			}

			treeNode.Collapse().ClearChildren()
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

	projectCreateModal = tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(alignment, 50, 1, true).
		AddItem(nil, 0, 1, false)

	return projectCreateModal
}

func createDirectory(parent *tview.TreeNode, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Invalid Collection")
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(path, name)

		node := CreateNode(name, path, entry.IsDir())

		child := tview.
			NewTreeNode(node.Name(true, false)).
			SetExpanded(false).
			SetReference(node)

		parent.AddChild(child)
	}
}
