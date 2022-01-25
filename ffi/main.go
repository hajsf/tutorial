package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "main.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(1, 2))

}

// no new lines entry allowed between the `#include` and the `import C`
// build tags will not work and will give error when working with cgo, below is not workable
// go run .
// OR (preferable)
// go run github.io/hajsf/ffi
// go run main.go will not work as it will consider main.go only, and not consider the C file
