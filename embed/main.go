package main

import (
	"fmt"
	"strings"

	_ "embed"
)

var Version string = strings.TrimSpace(version)

//go:embed version_prod.go
var src string // embed the source code of the version_prod.go

func main() {
	fmt.Printf("Version %q\n", Version)
	fmt.Print(src)
}

// $ go run .
// Version "dev"

// $ go run -tags prod .
// Version "0.0.1"
