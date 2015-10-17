package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/net/websocket"
)

var indexTmpl = template.Must(template.ParseFiles("assets/index.html"))

// StartServer starts http-server and servers frontend code
// for benchmark results display.
func StartServer(bind string, ch chan BenchmarkSet) error {
	http.HandleFunc("/", handler)
	http.Handle("/ws", websocket.Handler(wshandler))

	go StartBrowser("http://localhost" + bind)
	return http.ListenAndServe(bind, nil)
}

// handler handles index page.
func handler(w http.ResponseWriter, r *http.Request) {
	err := indexTmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// wshandler is a handler for websocket connection.
func wshandler(ws *websocket.Conn) {
	for {
		_, err := ws.Write([]byte("Some info from server"))
		fmt.Println("[DEBUG] WebSocket send", err)
		_ = err
		time.Sleep(1 * time.Second)
	}
}

// StartBrowser tries to open the URL in a browser
// and reports whether it succeeds.
//
// Orig code: golang.org/x/tools/cmd/cover/html.go
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
