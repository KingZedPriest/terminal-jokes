package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

var (
	app      *tview.Application
	textView *tview.TextView
)

func getAndDrawJoke() {
	// Fetch joke from the web

	//Update the UI with the joke
	textView.Clear()
	timeStr := fmt.Sprintf("[gray]%s", time.Now().Format(time.RFC1123))
	fmt.Fprintln(textView, timeStr)
}

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
