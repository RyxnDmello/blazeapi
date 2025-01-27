package controls

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Controls(app *tview.Application, query *tview.Flex, pages *tview.Pages, project *tview.TreeView, response *tview.Flex) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			CloseQueryBody(pages)
			app.SetFocus(query)

		case tcell.KeyCtrlP:
			CloseQueryBody(pages)
			app.SetFocus(project)

		case tcell.KeyCtrlR:
			CloseQueryBody(pages)
			app.SetFocus(response)

		case tcell.KeyCtrlB:
			if query.HasFocus() {
				ShowQueryBody(pages)
				break
			}

			ShowQuery(app, pages, query)

		case tcell.KeyEsc:
			name, _ := pages.GetFrontPage()

			if name == "QUERY_BODY_MODAL" {
				ShowQuery(app, pages, query)
				break
			}
		}

		return event
	})
}
