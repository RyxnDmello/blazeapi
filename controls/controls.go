package controls

import (
	"blazeapi/query"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var modals []string = []string{
	"CREATE_FOLDER_MODAL",
	"CREATE_FILE_MODAL",
	"DELETE_NODE_MODAL",
	"QUERY_BODY_MODAL",
}

func Controls(
	app *tview.Application,
	pages *tview.Pages,
	query *query.Query,
	queryLayout *tview.Flex,
	queryBodyModal *tview.Flex,
	project *tview.TreeView,
	createFileModal *tview.Flex,
	createFolderModal *tview.Flex,
	deleteNodeModal *tview.Flex,
	response *response.Response,
	responseLayout *tview.Flex,
) {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			CloseEveryModal(pages)
			app.SetFocus(queryLayout)

		case tcell.KeyCtrlB:
			if queryLayout.HasFocus() {
				OpenModal("QUERY_BODY_MODAL", pages)
				break
			}

			if IsOpen("QUERY_BODY_MODAL", pages) {
				CloseModal("QUERY_BODY_MODAL", pages)
				app.SetFocus(queryLayout)
			}

		case tcell.KeyCtrlP:
			CloseEveryModal(pages)
			app.SetFocus(project)

		case tcell.KeyCtrlA:
			if project.HasFocus() {
				OpenModal("CREATE_FILE_MODAL", pages)
				break
			}

			if IsOpen("CREATE_FILE_MODAL", pages) {
				CloseModal("CREATE_FILE_MODAL", pages)
				app.SetFocus(project)
			}

		case tcell.KeyCtrlF:
			if project.HasFocus() {
				OpenModal("CREATE_FOLDER_MODAL", pages)
				break
			}

			if IsOpen("CREATE_FOLDER_MODAL", pages) {
				CloseModal("CREATE_FOLDER_MODAL", pages)
				app.SetFocus(project)
			}

		case tcell.KeyCtrlD:
			if project.HasFocus() {
				OpenModal("DELETE_NODE_MODAL", pages)
				break
			}

			if IsOpen("DELETE_NODE_MODAL", pages) {
				CloseModal("DELETE_NODE_MODAL", pages)
				app.SetFocus(project)
			}

		case tcell.KeyCtrlR:
			CloseEveryModal(pages)
			app.SetFocus(responseLayout)

		case tcell.KeyEsc:
			Escape(app, modals, pages)
		}

		return event
	})
}
