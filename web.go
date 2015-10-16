package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

// StartServer starts http-server and servers frontend code
// for benchmark results display.
func StartServer(bind string, ch chan BenchmarkSet) error {
	http.HandleFunc("/", handler)
	go StartBrowser("http://localhost" + bind)
	return http.ListenAndServe(bind, nil)
}

// handler handles index page.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
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
