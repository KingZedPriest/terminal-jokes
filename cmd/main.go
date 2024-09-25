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
	app         *tview.Application
	textView    *tview.TextView
	jokeRefresh *time.Ticker
)

type Payload struct {
	Value string `json:"value"`
}

func getAndDrawJoke() {
	result, err := http.Get("https://api.chucknorris.io/jokes/random?category=science")
	if err != nil {
		textView.SetText(fmt.Sprintf("Error fetching joke: %v", err))
		app.Draw()
		return
	}

	payloadBytes, err := io.ReadAll(result.Body)
	if err != nil {
		textView.SetText(fmt.Sprintf("Error reading joke data: %v", err))
		app.Draw()
		return
	}

	payload := &Payload{}
	err = json.Unmarshal(payloadBytes, payload)
	if err != nil {
		textView.SetText(fmt.Sprintf("Error parsing joke data: %v", err))
		app.Draw()
		return
	}

	textView.Clear()
	fmt.Fprintln(textView, payload.Value)
	timeStr := fmt.Sprintf("\n\n[gray]%s", time.Now().Format(time.RFC1123))
	fmt.Fprintln(textView, timeStr)
}

func refreshJoke() {
	jokeRefresh = time.NewTicker(time.Second * 30)
	for {
		select {
		case <-jokeRefresh.C:
			getAndDrawJoke()
			app.Draw()
		}
	}
}

func main() {
	app = tview.NewApplication()

	// Define colors
	textColor := tcell.ColorWhite
	headerColor := tcell.ColorLightCyan
	footerColor := tcell.ColorGray

	// Create text views
	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("[::b][white]Chuck Norris Joke Teller!  [::-]").
		SetTextColor(headerColor).
		SetDynamicColors(true)

	textView = tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(textColor)

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("[::b][gray]A Chuck Norris joke app for your terminal. [::-]").
		SetTextColor(footerColor).
		SetDynamicColors(true)

	// Create grid
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(header, 0, 0, 1, 3, 0, 0, false).
		AddItem(textView, 1, 0, 1, 3, 0, 0, false).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	// Start joke fetching and refreshing
	getAndDrawJoke()
	go refreshJoke()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
