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
	var sockets Sockets
	http.Handle("/sse", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		socket := r.RemoteAddr
		sockets = append(sockets, socket)
		fmt.Println("sockets connected", sockets)
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Internal error", 500)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		setupCORS(&w, r)

		// send custom events
		// Event name: join // used in the client as: eventSource.addEventListener("join", handleReceiveMessage);
		// Event data: %v => data // used in the client as: const eventData = JSON.parse(event.data);
		// Event id: nid:500 // used in the client as: eventSource.lastEventId
		//fmt.Fprintf(w, "event: join\ndata: %v\nid:500\n\n", data)
		// Important note: \n does a line break. \n\n means the end of the message, do not forget to add that.
		fmt.Fprintf(w, "event: join\ndata: Welcome %v\n\n", socket)
		flusher.Flush()

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
				fmt.Println(len(sockets))
				sockets.RemoveElement(socket)
				fmt.Println(len(sockets))
				fmt.Println("sockets connected", sockets)
				fmt.Println(len(sockets))
				return
			}
		}
	}))

	// Retry each 5 seconds
	// fmt.Fprintf(w, `retry:5000\ndata: ${JSON.stringify(message)}\n\n`);

	if err := http.ListenAndServe(":1235", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Cache-Control", "no-cache")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
