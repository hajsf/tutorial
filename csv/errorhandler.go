package main

import "log"

func checkError(message string, err error) {
	// Error Logging
	if err != nil {
		log.Fatal(message, err)
	}
}
