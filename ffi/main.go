package main

// #include "main.h"
// #include  <stdio.h>
// int Add2(int x)
// {
//	    printf("Welcome from inline C function\n");
//	    return x + 2;
// }
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(1, 2))
	fmt.Println(C.Add2(5))
}

// no new lines entry allowed between the `#include` and the `import C`
// build tags will not work and will give error when working with cgo, below is not workable
// It is possible to pass some additional info to cgo via special comments // #cgo
// https://pkg.go.dev/cmd/cgo
// go run .
// OR (preferable)
// go run github.io/hajsf/ffi
// go run main.go will not work as it will consider main.go only, and not consider the C file
