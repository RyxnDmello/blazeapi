package cmd

import (
	"blazeapi/controls"
	QUERY "blazeapi/query"

	"github.com/rivo/tview"
)

var (
	query    *tview.Flex
	project  *tview.TreeView
	response *tview.TextView
)

func Blaze(app *tview.Application) *tview.Pages {
	project = tview.NewTreeView()
	project.SetBorder(true)

	response = tview.NewTextView()
	response.SetBorder(true)

	query, queryBody := QUERY.InitializeQuery(app, response)

	grid := tview.
		NewGrid().
		SetRows(3, 0).
		AddItem(query, 0, 0, 1, 2, 0, 0, true).
		AddItem(project, 1, 0, 1, 1, 0, 0, false).
		AddItem(response, 1, 1, 1, 1, 0, 0, false)

	pages := tview.
		NewPages().
		AddPage("QUERY_BODY_MODAL", queryBody, true, false).
		AddPage("MAIN", grid, true, true)

	controls.Controls(app, pages, query)

	return pages
}
