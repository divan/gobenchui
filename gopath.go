package main

import (
	"fmt"
	"os"
	"strings"
)

// GOPATH extracts first gopath from your env variable GOPATH.
func GOPATH() string {
	gopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		fmt.Fprintf(os.Stderr, "GOPATH not set, aborting")
		os.Exit(1)
	}
	paths := strings.Split(gopath, ":")
	return paths[0]
}
