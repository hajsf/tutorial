package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.ListenAndServe(":8888", nil)
}

//go:embed static
var embededFiles embed.FS

// Files started with . or _ are not embeded

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode, open: http://localhost:8888/index.html")
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

// $ go run github.io/hajsf/server
// open: http://localhost:8888/index.html
