package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	if err := app.SetRoot(nil, true).Run(); err != nil {
		panic(err)
	}
}
