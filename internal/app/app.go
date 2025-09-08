package app

import (
	"github.com/rivo/tview"

	"github.com/ryxndmello/flame/internal/explorer"
	"github.com/ryxndmello/flame/internal/query"
	"github.com/ryxndmello/flame/internal/response"
)

var app *tview.Application = tview.NewApplication()

func Run() error {
	explorerView := explorer.
		NewDefaultExplorer(app)

	responseView := response.
		NewDefaultResponse(app)

	queryView := query.
		NewDefaultQuery(app, responseView)

	middleLayout := tview.
		NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(explorerView.GetTree(), 40, 1, true).
		AddItem(responseView.GetLayout(), 0, 1, false)

	layout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(queryView.GetLayout(), 1, 1, true).
		AddItem(middleLayout, 0, 1, true)

	pages := tview.
		NewPages().
		AddPage("MAIN", layout, true, true)

	return app.SetRoot(pages, true).Run()
}
