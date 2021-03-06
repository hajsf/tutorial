package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Text: ")
	scanner.Scan()
	fmt.Println(scanner.Text())

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	/*
		for {
			fmt.Print("Enter Text: ")
			// reads user input until \n by default
			scanner.Scan()
			// Holds the string that was scanned
			text := scanner.Text()
			if len(text) != 0 {
				fmt.Println(text)
			} else {
				// exit if user entered an empty string
				break
			}

		} */

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}
