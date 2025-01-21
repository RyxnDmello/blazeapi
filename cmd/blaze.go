package cmd

import (
	QUERY "blazeapi/query"

	"github.com/rivo/tview"
)

var (
	query    *tview.Flex
	project  *tview.TreeView
	response *tview.TextView
)

func Blaze(app *tview.Application) *tview.Grid {
	project = tview.NewTreeView()
	project.SetBorder(true)

	response = tview.NewTextView()
	response.SetBorder(true)

	query = QUERY.InitializeQuery(app, response)

	grid := tview.
		NewGrid().
		SetRows(3, 0).
		AddItem(query, 0, 0, 1, 2, 0, 0, true).
		AddItem(project, 1, 0, 1, 1, 0, 0, true).
		AddItem(response, 1, 1, 1, 1, 0, 0, true)

	return grid
}
