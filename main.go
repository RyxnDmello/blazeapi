package main

import (
	"blazeapi/cmd"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {

	layout := cmd.Blaze(app)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
