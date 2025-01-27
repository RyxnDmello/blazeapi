package controls

import "github.com/rivo/tview"

func ShowQuery(app *tview.Application, pages *tview.Pages, query *tview.Flex) {
	CloseQueryBody(pages)
	app.SetFocus(query)
}

func ShowQueryBody(pages *tview.Pages) {
	pages.SendToFront("QUERY_BODY_MODAL").ShowPage("QUERY_BODY_MODAL")
}

func CloseQueryBody(pages *tview.Pages) {
	pages.SendToBack("QUERY_BODY_MODAL").HidePage("QUERY_BODY_MODAL")
}
