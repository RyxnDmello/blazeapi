package project

import "github.com/rivo/tview"

func InitializeProject(app *tview.Application) *tview.TreeView {
	project := tview.
		NewTreeView()

	project.
		SetBorder(true).
		SetTitle("  Manager ").
		SetTitleAlign(tview.AlignLeft)

	return project
}
