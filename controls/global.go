package controls

import (
	"blazeapi/utils"
	"strings"

	"github.com/rivo/tview"
)

func OpenModal(name string, pages *tview.Pages) {
	CloseEveryModal(pages)

	if strings.Contains(name, "MODAL") {
		pages.SendToFront(name).ShowPage(name)
	}
}

func CloseModal(name string, pages *tview.Pages) {
	if !strings.Contains(name, "MODAL") {
		return
	}

	pages.SendToBack(name).HidePage(name)
}

func CloseEveryModal(pages *tview.Pages) {
	names := pages.GetPageNames(false)

	for _, name := range names {
		if !strings.Contains(name, "MODAL") {
			continue
		}

		pages.SendToBack(name).HidePage(name)
	}
}

func IsOpen(name string, pages *tview.Pages) (open bool) {
	front, _ := pages.GetFrontPage()
	return front == name
}

func Escape(app *tview.Application, pages *tview.Pages) {
	if IsOpen("QUERY_BODY_MODAL", pages) {
		CloseModal("QUERY_BODY_MODAL", pages)
		return
	}

	if IsOpen("PROJECT_CREATE_NODE_MODAL", pages) {
		CloseModal("PROJECT_CREATE_NODE_MODAL", pages)
		return
	}

	utils.Exit(app, "Terminated Successfully")
}
