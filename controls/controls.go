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
	queryBodyModal *tview.Flex,
	project *tview.TreeView,
	projectCreateModal *tview.Flex,
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
				OpenModal("QUERY_BODY_MODAL", pages)
				break
			}

			if ModalOpen("QUERY_BODY_MODAL", pages) {
				CloseModal("QUERY_BODY_MODAL", pages)
				app.SetFocus(queryLayout)
			}

		case tcell.KeyCtrlP:
			CloseEveryModal(pages)
			app.SetFocus(project)

		case tcell.KeyCtrlN:
			if project.HasFocus() {
				OpenModal("PROJECT_CREATE_MODAL", pages)
				break
			}

			if ModalOpen("PROJECT_CREATE_MODAL", pages) {
				CloseModal("PROJECT_CREATE_MODAL", pages)
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
