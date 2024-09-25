package main

import (
	"github.com/rivo/tview"
)

var (
	app *tview.Application
)

func main() {
	app = tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText("First Go App")

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
