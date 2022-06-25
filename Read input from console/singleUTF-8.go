package main

import (
	"bufio"
	"fmt"
	"os"
)

func main4() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	// prints the unicode code point of the character
	fmt.Println(char)
}
