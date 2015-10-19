package main

import (
	"fmt"
	"os"
	"sync"
)

// Benchmaker represents tool used for running go benchmarks.
type Benchmaker interface {
	Benchmark(workspace Workspace) (string, error)
}

var once sync.Once

// Gotool is a default 'go test' tool.
type GoTool struct {
}

// Benchmark runs go benchmarks. Implements Benchmarker.
func (GoTool) Benchmark(workspace *Workspace, benchRegexp string) (string, error) {
	// as we used cloned workspace outside of original GOPATH,
	// add current temp directory to GOPATH variable in order to
	// GO15VENDOREXPERIMENT work correctly.
	once.Do(func() {
		if v, ok := os.LookupEnv("GO15VENDOREXPERIMENT"); ok && v == "1" {
			gopath := os.Getenv("GOPATH")
			gopath = fmt.Sprintf("%s:%s", gopath, workspace.Gopath())
			fmt.Println("[INFO] Detected GO15VENDOREXPERIMENT, setting GOPATH to", gopath)
			if err := os.Setenv("GOPATH", gopath); err != nil {
				fmt.Println("[ERROR] cannot update GOPATH for benchmark")
				return
			}
		}
	})

	return Run(workspace.Path(), "go", "test", "-run", "XXXXXX", "-bench", benchRegexp, "-benchmem")
}
