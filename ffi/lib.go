//go:build lib
// +build lib

package main

import "C"
import "fmt"

//export HelloWorld
func HelloWorld() {
	fmt.Printf("hello world")
}

func main() {}

// go build -tags lib -buildmode=c-shared -o golib.a lib.go
