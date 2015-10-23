package main

import (
	"fmt"
	"os"
	"strings"
)

// GOPATH extracts first gopath from your env variable GOPATH.
func GOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Fprintf(os.Stderr, "GOPATH not set, aborting")
		os.Exit(1)
	}
	paths := strings.Split(gopath, string(os.PathListSeparator))
	return paths[0]
}
