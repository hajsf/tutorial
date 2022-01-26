package main

//#include  <stdio.h>
//int Add(int a, int b){
//    printf("Welcome from external C function\n");
//    return a + b;
//}
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(5, 2))
}

// no new lines entry allowed between the `#include` and the `import C`
// build tags will not work and will give error when working with cgo, below is not workable
// It is possible to pass some additional info to cgo via special comments // #cgo
// https://pkg.go.dev/cmd/cgo
// go run .
// OR (preferable)
// go run github.io/hajsf/ffi
// go run main.go will work also as every thing in the same file

// go build main.go
