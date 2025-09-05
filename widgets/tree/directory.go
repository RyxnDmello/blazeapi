package tree

import (
	"os"
	"path"

	"github.com/rivo/tview"
)

func AddDirectory(treeNode *tview.TreeNode) {
	node, ok := treeNode.GetReference().(*Node)

	if !ok || !node.IsDir() {
		return
	}

	label := node.
		GetDisplay(true, false, true)

	treeNode.
		SetText(label)

	entries := read(node.path)

	for _, entry := range entries {
		path := path.Join(node.path, entry.Name())

		newNode := NewNode().
			SetPath(path).
			SetParent(node).
			SetName(entry.Name()).
			SetDir(entry.IsDir()).
			SetLevel(node.level + 1).
			SetPadding(node.padding).
			SetSelectedIcon('\uf15b').
			SetUnselectedIcon('\uf15b').
			SetSelectedStyle(node.GetSelectedStyleDecompose()).
			SetUnselectedStyle(node.GetUnselectedStyleDecompose())

		if entry.IsDir() {
			newNode.
				SetSelectedIcon('\uf07c').
				SetUnselectedIcon('\uf07b')
		}

		treeNode.
			AddChild(newNode.Draw())
	}
}

func Toggle(node *Node, treeNode *tview.TreeNode) {
	children := treeNode.
		GetChildren()

	if len(children) == 0 {
		AddDirectory(treeNode)
	}

	expanded := !treeNode.
		IsExpanded()

	label := node.
		GetDisplay(expanded, !expanded, true)

	treeNode.
		SetText(label).
		SetExpanded(expanded)
}

func read(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Unable To Read The Provided Path")
	}

	return sort(entries)
}

func sort(entries []os.DirEntry) []os.DirEntry {
	files := make([]os.DirEntry, 0)
	folders := make([]os.DirEntry, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry)
			continue
		}

		files = append(files, entry)
	}

	return append(folders, files...)
}
