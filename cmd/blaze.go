package cmd

import (
	CONTROLS "blazeapi/controls"
	PROJECT "blazeapi/project"
	QUERY "blazeapi/query"
	RESPONSE "blazeapi/response"

	"github.com/rivo/tview"
)

var (
	query       QUERY.Query
	queryLayout *tview.Flex
	bodyModal   *tview.Flex

	project         *tview.TreeView
	createNodeModal *tview.Flex

	response       RESPONSE.Response
	responseLayout *tview.Flex
)

func Blaze(app *tview.Application) *tview.Pages {
	response, responseLayout = RESPONSE.InitializeResponse(app)
	query, queryLayout, bodyModal = QUERY.InitializeQuery(app, response)
	project, createNodeModal = PROJECT.InitializeProject(app, query, response)

	layout := tview.
		NewFlex().
		AddItem(project, 0, 1, true).
		AddItem(responseLayout, 0, 2, true)

	main := tview.
		NewGrid().
		SetRows(3, 0).
		AddItem(queryLayout, 0, 0, 1, 2, 0, 0, false).
		AddItem(layout, 1, 0, 1, 2, 0, 0, true)

	pages := tview.
		NewPages().
		AddPage("BODY_MODAL", bodyModal, true, false).
		AddPage("CREATE_NODE_MODAL", createNodeModal, true, false).
		AddPage("MAIN", main, true, true)

	CONTROLS.Controls(
		app,
		pages,
		query,
		queryLayout,
		bodyModal,
		project,
		createNodeModal,
		response,
		responseLayout,
	)

	return pages
}
