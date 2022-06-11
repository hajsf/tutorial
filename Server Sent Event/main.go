package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	counter int
)

func main() {
	http.Handle("/sse/dashboard", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from here")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Internal error", 500)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		setupCORS(&w, r)

		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				mu.Lock()
				counter++
				c := counter
				mu.Unlock()
				fmt.Fprintf(w, "data: %v\n\n", c)
				flusher.Flush()
			case <-r.Context().Done():
				return
			}
		}
	}))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Cache-Control", "no-cache")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
