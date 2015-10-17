//go:generate go-bindata-assetfs assets/...
package main

import (
	"encoding/json"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"time"

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
		wshandler(ws, resCh, runCh, info)
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

// WSData used for WebSocket commincation with frontend.
type WSData struct {
	Type   string       `json:"type"` // "status" or "result"
	Result BenchmarkSet `json:"result,omitempty"`

	Status    Status    `json:"status,omitempty"`
	Progress  float64   `json:"progress,omitempty"`
	Commit    Commit    `json:"commit,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

// wshandler is a handler for websocket connection.
func wshandler(ws *websocket.Conn, resCh chan BenchmarkSet, runCh chan BenchmarkRun, info *Info) {
	defer func() {
		fmt.Println("[DEBUG] Closing connection")
		ws.Close()
	}()
	for {
		select {
		case status, ok := <-runCh:
			if !ok {
				runCh = nil
				if resCh == nil {
					return
				}
			}
			data := WSData{
				Type:      "status",
				Status:    InProgress,
				Commit:    status.Commit,
				StartTime: status.StartTime,
			}

			// Ignore error, as we can't help much here
			_ = sendJSON(ws, data)
		case set, ok := <-resCh:
			if !ok {
				resCh = nil
				if runCh == nil {
					return
				}
			}
			data := WSData{
				Type:   "result",
				Result: set,
				Status: InProgress,
			}

			// Ignore error, as we can't help much here
			_ = sendJSON(ws, data)
		}
	}
}

// sendJSON is a wrapper for sending JSON encoded data to websocket
func sendJSON(ws *websocket.Conn, data interface{}) error {
	body, err := json.MarshalIndent(data, "  ", "    ")
	if err != nil {
		fmt.Println("[ERROR] JSON encoding failed", err)
		return err
	}

	_, err = ws.Write(body)
	if err != nil {
		fmt.Println("[ERROR] WebSocket send failed", err)
		return err
	}

	return nil
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
