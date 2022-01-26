package main

//#cgo CFLAGS: -g -Wall
//#cgo LDFLAGS: -L. -ladd
//#include "libadd.h"
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
// go run main.go will work also as the library is linked in the same file

// To run the executable binary, both executable binar and the dynamic library should be in the same folder
