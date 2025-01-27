package project

import "github.com/rivo/tview"

func InitializeProject() *tview.TreeView {
	project := tview.
		NewTreeView()

	project.
		SetBorder(true).
		SetTitle("  Manager ").
		SetTitleAlign(tview.AlignLeft)

	return project
}
