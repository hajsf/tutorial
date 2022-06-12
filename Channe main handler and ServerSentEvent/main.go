package main

import (
	"fmt"
	"net/http"
	"time"
)

type DataPasser struct {
	logs chan string
}

func main() {
	passer := &DataPasser{logs: make(chan string)}
	t := time.NewTicker(time.Second)
	defer t.Stop()

	go func() {
		for range t.C {
			passer.logs <- time.Now().String()
		}
	}()

	http.HandleFunc("/", passer.handleHello)
	http.ListenAndServe(":9999", nil)

}

func (p *DataPasser) handleHello(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	w.Header().Set("Content-Type", "text/event-stream")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Internal error", 500)
		return
	}
	flusher.Flush()
	done := r.Context().Done()
	defer fmt.Println("EXIT")
	for {
		select {
		case <-done:
			// the client disconnected
			fmt.Println("Client disconnected")
			return
		case m := <-p.logs:
			if _, err := fmt.Fprintf(w, "data: %s\n\n", m); err != nil {
				// Write to connection failed. Subsequent writes will probably fail.
				fmt.Println("Writting to socket failed")
				return
			}
			flusher.Flush()
		}
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Cache-Control", "no-cache")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
