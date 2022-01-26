//go:build lib
// +build lib

package main

import "C"
import "fmt"

//export HelloWorld
func HelloWorld() {
	fmt.Printf("hello world fro GO\n")
}

func main() {}

// go build -tags lib -buildmode=c-shared -o ../libraries/libgo.a lib.go
