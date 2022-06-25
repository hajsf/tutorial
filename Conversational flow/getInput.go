package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(msg string) (input string) {
	fmt.Print(msg, " ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	// fmt.Println(input)
	return
}
