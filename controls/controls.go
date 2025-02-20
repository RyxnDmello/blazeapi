package controls

import (
	"blazeapi/query"
	"blazeapi/response"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Controls(
	app *tview.Application,
	pages *tview.Pages,
	query query.Query,
	queryLayout *tview.Flex,
	bodyModal *tview.Flex,
	project *tview.TreeView,
	createNodeModal *tview.Flex,
	deleteNodeModal *tview.Flex,
	response response.Response,
	responseLayout *tview.Flex,
) {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			CloseEveryModal(pages)
			app.SetFocus(queryLayout)

		case tcell.KeyCtrlB:
			if queryLayout.HasFocus() {
				OpenModal("BODY_MODAL", pages)
				break
			}

			if IsOpen("BODY_MODAL", pages) {
				CloseModal("BODY_MODAL", pages)
				app.SetFocus(queryLayout)
			}

		case tcell.KeyCtrlP:
			CloseEveryModal(pages)
			app.SetFocus(project)

		case tcell.KeyCtrlN:
			if project.HasFocus() {
				OpenModal("CREATE_NODE_MODAL", pages)
				break
			}

			if IsOpen("CREATE_NODE_MODAL", pages) {
				CloseModal("CREATE_NODE_MODAL", pages)
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
			Escape(app, pages)
		}

		return event
	})
}
