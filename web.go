package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"

	"golang.org/x/net/websocket"
)

var indexTmpl = template.Must(template.ParseFiles("assets/index.html"))

// StartServer starts http-server and servers frontend code
// for benchmark results display.
func StartServer(bind string, ch chan BenchmarkSet) error {
	http.HandleFunc("/", handler)
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		wshandler(ws, ch)
	}))

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
func wshandler(ws *websocket.Conn, ch chan BenchmarkSet) {
	for set := range ch {
		data, err := json.MarshalIndent(set, "  ", "    ")
		if err != nil {
			fmt.Println("[ERROR] JSON encoding failed", err)
			continue
		}

		_, err = ws.Write(data)
		if err != nil {
			fmt.Println("[ERROR] WebSocket send failed", err)
			continue
		}
	}

	fmt.Println("[DEBUG] Closing connection")
	ws.Close()
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
