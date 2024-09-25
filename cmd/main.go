package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app      *tview.Application
	textView *tview.TextView
)

type Payload struct {
	Value string
}

func getAndDrawJoke() {
	// Fetch joke from the web
	result, err := http.Get("https://api.chucknorris.io/jokes/random?category=science")
	if err != nil {
		panic(err)
	}

	payloadBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	payload := &Payload{}
	err = json.Unmarshal(payloadBytes, payload)
	if err != nil {
		panic(err)
	}

	//Update the UI with the joke
	textView.Clear()
	fmt.Fprintln(textView, payload.Value)
	timeStr := fmt.Sprintf("\n\n[gray]%s", time.Now().Format(time.RFC1123))
	fmt.Fprintln(textView, timeStr)
}

func refreshJoke() {
	tick := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-tick.C:
			getAndDrawJoke()
			app.Draw()
		}
	}
}

func main() {
	app = tview.NewApplication()

	textView = tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLime)

	go refreshJoke()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
