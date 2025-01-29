package controls

import (
	"blazeapi/utils"

	"github.com/rivo/tview"
)

func Escape(app *tview.Application, pages *tview.Pages, query *tview.Flex) {
	name, _ := pages.GetFrontPage()

	if name == "QUERY_BODY_MODAL" {
		ShowQuery(app, pages, query)
		return
	}

	utils.Exit(app, "Terminated Successfully")
}
