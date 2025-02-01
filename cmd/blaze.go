package cmd

import (
	CONTROLS "blazeapi/controls"
	PROJECT "blazeapi/project"
	QUERY "blazeapi/query"
	RESPONSE "blazeapi/response"

	"github.com/rivo/tview"
)

var (
	query          QUERY.Query
	queryLayout    *tview.Flex
	queryBodyModal *tview.Flex

	project *tview.TreeView

	response       RESPONSE.Response
	responseLayout *tview.Flex
)

func Blaze(app *tview.Application) *tview.Pages {
	response, responseLayout = RESPONSE.InitializeResponse(app)
	query, queryLayout, queryBodyModal = QUERY.InitializeQuery(app, response)
	project = PROJECT.InitializeProject(app, query, response)

	main := tview.
		NewGrid().
		SetRows(3, 0).
		AddItem(queryLayout, 0, 0, 1, 2, 0, 0, false).
		AddItem(project, 1, 0, 1, 1, 0, 0, true).
		AddItem(responseLayout, 1, 1, 1, 1, 0, 0, false)

	pages := tview.
		NewPages().
		AddPage("QUERY_BODY_MODAL", queryBodyModal, true, false).
		AddPage("MAIN", main, true, true)

	CONTROLS.Controls(app, pages, queryLayout, project, responseLayout)

	return pages
}
