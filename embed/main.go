package main

import (
	"fmt"
	"strings"

	_ "embed" // the _ used to enforce import, as GO compiler will understand there is unused import,
)

var Version string = strings.TrimSpace(version)

//go:embed version_prod.go
var src string // embed the source code of the version_prod.go
// if the embeded item is a folder, then files started with . or _ are not embeded

func main() {
	fmt.Printf("Version %q\n", Version)
	fmt.Print(src)
}

// $ go run .
// Version "dev"

// $ go run -tags prod .
// Version "0.0.1"
