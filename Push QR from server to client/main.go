package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func main() {
	passer := &DataPasser{logs: make(chan string)}

	http.HandleFunc("/sse/dashboard", passer.handleHello)
	go http.ListenAndServe(":1234", nil)
	/*
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					//	fmt.Println("Tick at", t)
					// passer.logs <- buffer.String()
				}
			}
		}()
	*/
	store.DeviceProps.Os = proto.String("WhatsApp GO")
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:datastore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
	GetQR:
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}

		for evt := range qrChan {
			switch evt.Event {
			case "success":
				{
					passer.logs <- "success"
					fmt.Println("Login event: success")
				}
			case "timeout":
				{
					passer.logs <- "timeout/Refreshing"
					fmt.Println("Login event: timeout")
					goto GetQR
				}
			case "code":
				{
					fmt.Println("new code recieved")
					fmt.Println(evt.Code)
					passer.logs <- evt.Code
				}
			}
		}
	} else {
		// Already logged in, just connect
		passer.logs <- "Already logged"
		fmt.Println("Already logged")
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
