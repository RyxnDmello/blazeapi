package project

import (
	"os"
	"path/filepath"

	"blazeapi/query"
	"blazeapi/response"

	"github.com/rivo/tview"
)

func InitializeProject(app *tview.Application, query query.Query, response response.Response) *tview.TreeView {
	root := tview.NewTreeNode("./test")

	createDirectory(root, "./test")

	project := tview.
		NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphics(false).
		SetTopLevel(1)

	project.SetSelectedFunc(func(treeNode *tview.TreeNode) {
		node, ok := treeNode.GetReference().(Node)

		if !ok {
			return
		}

		if !node.IsCollection() {
			api := CreateAPI(node.path)
			query.SetUrl(api.Url)
			return
		}

		if treeNode.IsExpanded() {
			treeNode.ClearChildren().SetExpanded(false)
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

	return project
}

func createDirectory(parent *tview.TreeNode, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Invalid Collection")
	}

	for i := 0; i < len(entries); i++ {
		name := entries[i].Name()
		path := filepath.Join(path, name)

		node := CreateNode(name, path, entries[i].IsDir())

		child := tview.
			NewTreeNode(node.Name(true)).
			SetExpanded(false).
			SetReference(node)

		parent.AddChild(child)
	}
}
