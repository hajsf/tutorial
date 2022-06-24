package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jchv/go-webview2" // go install github.com/jchv/go-webview2@latest
)

type IncrementResult struct {
	Count uint `json:"count"`
}

func main() {
	var count uint = 0

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Window:    nil,
		Debug:     false,
		DataPath:  "",
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			Width:  800,
			Height: 600,
			IconId: 2,
			Center: true},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()

	http.HandleFunc("/", index)

	go http.ListenAndServe(":1234", nil)

	w.SetTitle("Golang WebView 2 example")
	w.SetSize(800, 600, webview2.HintFixed)

	// load a local HTML file.
	c, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}
	w.Navigate(filepath.Join(c, "templates/index.html")) // w.Navigate("http://localhost:1234/")

	w.Bind("increment", func() IncrementResult {
		count++
		return IncrementResult{Count: count}
	})

	w.Init("alert('hi')")

	w.Run()
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
