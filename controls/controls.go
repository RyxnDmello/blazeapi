package controls

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Controls(app *tview.Application, pages *tview.Pages, query *tview.Flex) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlB:
			if query.HasFocus() {
				pages.SendToFront("QUERY_BODY_MODAL").ShowPage("QUERY_BODY_MODAL")
				break
			}

			pages.SendToFront("MAIN").HidePage("QUERY_BODY_MODAL")

		case tcell.KeyEsc:
			if name, _ := pages.GetFrontPage(); name == "QUERY_BODY_MODAL" {
				pages.SendToBack("QUERY_BODY_MODAL").HidePage("QUERY_BODY_MODAL")
				break
			}
		}

		return event
	})
}
