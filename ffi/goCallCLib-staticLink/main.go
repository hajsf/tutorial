package main

// #include "libadd.h"
import "C"
import "fmt"

func main() {
	x := C.Add(1, 2)
	fmt.Println(x)

}

// no new lines entry allowed between the `#include` and the `import C`
// build tags will not work and will give error when working with cgo, below is not workable
// It is possible to pass some additional info to cgo via special comments // #cgo
// https://pkg.go.dev/cmd/cgo
// go run .
// OR (preferable)
// go run github.io/hajsf/ffi
// go run main.go will not work as it will consider main.go only, and not consider the C file

// go build -o main .
// The lib will be integrated in the generated binary, so no need to have them together for running the app
