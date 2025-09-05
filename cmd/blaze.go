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

	primaryLayout := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(project, 40, 1, true).
		AddItem(responseLayout, 0, 1, true)

	layout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(queryLayout, 1, 1, false).
		AddItem(primaryLayout, 0, 1, true)

	pages := tview.
		NewPages().
		AddPage("QUERY_BODY_MODAL", queryBodyModal, true, false).
		AddPage("DELETE_NODE_MODAL", deleteNodeModal, true, false).
		AddPage("CREATE_FOLDER_MODAL", createFolderModal, true, false).
		AddPage("CREATE_FILE_MODAL", createFileModal, true, false).
		AddPage("MAIN", layout, true, true)

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
