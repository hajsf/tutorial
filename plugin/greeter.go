//go:build plug

package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

// exported as symbol named "Greeter"
var Greeter greeting

// go build -tags plug -buildmode=plugin -o greeter.so github.io/hajsf/plugin
