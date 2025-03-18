package cmd

import (
	CONTROLS "blazeapi/controls"
	PROJECT "blazeapi/project"
	QUERY "blazeapi/query"
	RESPONSE "blazeapi/response"

	"github.com/rivo/tview"
)

var (
	query             *QUERY.Query
	queryLayout       *tview.Flex
	queryBodyModal    *tview.Flex
	project           *tview.TreeView
	createFileModal   *tview.Flex
	createFolderModal *tview.Flex
	deleteNodeModal   *tview.Flex
	response          *RESPONSE.Response
	responseLayout    *tview.Flex
)

func Blaze(app *tview.Application) *tview.Pages {
	response, responseLayout = RESPONSE.InitializeResponse(app)
	query, queryLayout, queryBodyModal = QUERY.InitializeQuery(app, response)
	project, createFileModal, createFolderModal, deleteNodeModal = PROJECT.InitializeProject(app, query, response)

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
		AddPage("QUERY_BODY_MODAL", queryBodyModal, true, false).
		AddPage("DELETE_NODE_MODAL", deleteNodeModal, true, false).
		AddPage("CREATE_FOLDER_MODAL", createFolderModal, true, false).
		AddPage("CREATE_FILE_MODAL", createFileModal, true, false).
		AddPage("MAIN", main, true, true)

	CONTROLS.Controls(
		app,
		pages,
		query,
		queryLayout,
		queryBodyModal,
		project,
		createFileModal,
		createFolderModal,
		deleteNodeModal,
		response,
		responseLayout,
	)

	return pages
}
