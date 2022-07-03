package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	//	"github.com/zserge/lorca"

	"github.com/jchv/go-webview2"
)

func init() {
	//
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

//go:embed static
var staticFiles embed.FS

func main() {
	tmpDir := os.TempDir() + "\\cvsText"
	if p, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err = os.Mkdir(tmpDir, 0755)
		defer os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Printf("err 2: %v", err)
		} else {
			fmt.Println("temp created at:", p)
			_, exists := os.LookupEnv("cv")
			if !exists {
				//
				//err = os.Setenv(`cv`, tmpDir)
				_ = exec.Command(`SETX`, `cv`, tmpDir).Run()
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				//	fmt.Println("tmpDir: ", tmpDir) */
			} else {
				fmt.Println("Env exisit")
			}
		}
	} else {
		fmt.Println("checking Env ")
		_, exists := os.LookupEnv("cv")
		if !exists {
			//
			//err = os.Setenv(`cv`, tmpDir)
			_ = exec.Command(`SETX`, `cv`, tmpDir).Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Println("Env created")
			}
			//	fmt.Println("tmpDir: ", tmpDir) */
		} else {
			fmt.Println("Env exisit")
		}
	}
	go func() {
		// http.FS can be used to create a http Filesystem
		var staticFS = http.FS(staticFiles)
		fs := http.FileServer(staticFS) // embeded static files
		// Serve static files, to be embedded in the binary
		http.Handle("/static/", fs)

		http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
			http.ServeFile(rw, r, "http://localhost:3000/static/favicon.ico")
		})

		//	www := http.FileServer(http.Dir("/files/")) // side static files
		// Serve public files, to be beside binary
		http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./files"))))

		http.Handle("/pdf/", http.StripPrefix("/pdf/", http.FileServer(http.Dir(tmpDir))))

		//	defer os.RemoveAll(tempDir)
		/*	tmpDir := os.TempDir() + "\\scanner"
			if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
				err = os.Mkdir(tmpDir, 0755)
				if err != nil {
					fmt.Printf("err 2: %v", err)
				} else {
					fmt.Println(tmpDir)
					http.Handle("/pdf/", http.StripPrefix("/pdf/", http.FileServer(http.Dir(tmpDir))))
				}
			}  else {
				 fmt.Printf("\ntmpDir: %v already exixted", tmpDir)
			} */

		http.HandleFunc("/getSkills", getSkills)

		http.ListenAndServe(":3000", nil)

	}()
	// Start UI
	date := "2022-10-30"
	dateStamp, _ := time.Parse(layoutISO, date)
	today := time.Now()
	var url string
	if today.After(dateStamp) {
		url = "http://localhost:3000/static/expired.html"
	} else {
		url = "http://localhost:3000/static/"
	}

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Window:    nil,
		Debug:     false,
		DataPath:  "",
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			Width:  800,
			Height: 600,
			IconId: 2,
			Center: true},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()

	w.SetTitle("Golang WebView 2 example")
	w.SetSize(800, 600, webview2.HintFixed)

	w.Navigate(url)

	//	ui, err := lorca.New(url, "", 1200, 800)
	//	if err != nil {
	//		fmt.Println("error:", err)
	//	}
	//	defer ui.Close()

	// Bind Go function to be available in JS. Go function may be long-running and
	// blocking - in JS it's represented with a Promise.
	// ui.Bind("add", func(a, b int) int { return a + b })
	w.Bind("add", func(a, b int) int { return a + b })
	w.Run()

	// Call JS function from Go. Functions may be asynchronous, i.e. return promises
	//n := ui.Eval(`Math.random()`).Float()
	//fmt.Println(n)

	//n := w.Eval(`Math.random()`).Float()

	// Call JS that calls Go and so on and so on...
	//m := ui.Eval(`add(2, 3)`).Int()
	//fmt.Println(m)

	// Wait for the browser window to be closed
	//<-ui.Done()
}

/*
To return JSON to the client instead of text
	data := SomeStruct{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
*/
