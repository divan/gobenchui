//go:generate go-bindata-assetfs assets/...
package main

import (
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"

	"golang.org/x/net/websocket"
)

// indexTmpl is a html template for index page.
var indexTmpl *template.Template

func init() {
	indexTmpl = prepareTemplate()
}

// StartServer starts http-server and servers frontend code
// for benchmark results display.
func StartServer(bind string, resCh chan BenchmarkSet, runCh chan BenchmarkRun, info *Info) error {
	// Handle static files
	var fs http.FileSystem
	if DevMode() {
		fs = http.Dir("assets")
	} else {
		fs = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "assets"}
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(fs)))

	// Index page handler
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, info)
	}))

	// Websocket handler
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		wshandler(ws, resCh, runCh)
	}))

	go StartBrowser("http://localhost" + bind)

	return http.ListenAndServe(bind, nil)
}

// handler handles index page.
func handler(w http.ResponseWriter, r *http.Request, info *Info) {
	err := indexTmpl.Execute(w, info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[ERROR] failed to render template:", err)
		return
	}
}

// StartBrowser tries to open the URL in a browser
// and reports whether it succeeds.
//
// Orig. code: golang.org/x/tools/cmd/cover/html.go
func StartBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

// funcs for index template
var funcs = template.FuncMap{
	"last": func(a interface{}) interface{} {
		v := reflect.ValueOf(a)
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			return v.Index(v.Len() - 1).Interface()
		default:
			return nil
		}
	},
}

// DevMode returns true if app is running in development mode.
func DevMode() bool {
	devMode := os.Getenv("GOBENCHUI_DEV")
	return devMode != ""
}

// prepareTemplate prepares and parses template.
func prepareTemplate() *template.Template {
	t := template.New("index.html").Funcs(funcs)

	// read from local filesystem for development
	if DevMode() {
		return template.Must(t.ParseFiles("assets/index.html"))
	}

	data, err := Asset("assets/index.html")
	if err != nil {
		panic(err)
	}
	return template.Must(t.Parse(string(data)))
}
