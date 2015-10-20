package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Benchmarker represents tool used for running go benchmarks.
type Benchmarker interface {
	Benchmark(workspace *Workspace, regexp string) (string, error)
	Check(workspace *Workspace) bool
	Name() string
}

var once sync.Once

// as we used cloned workspace outside of original GOPATH,
// add current temp directory to GOPATH variable in order to
// GO15VENDOREXPERIMENT work correctly.
func honourVendorExperiment(workspace *Workspace) {
	once.Do(func() {
		if v := os.Getenv("GO15VENDOREXPERIMENT"); v == "1" {
			gopath := os.Getenv("GOPATH")
			gopath = fmt.Sprintf("%s:%s", gopath, workspace.Gopath())
			fmt.Println("[INFO] Detected GO15VENDOREXPERIMENT, setting GOPATH to", gopath)
			if err := os.Setenv("GOPATH", gopath); err != nil {
				fmt.Println("[ERROR] cannot update GOPATH for benchmark")
				return
			}
		}
	})
}

// GoTool is a default 'go test' tool.
type GoTool struct{}

// Name returns tool name. Implements Benchmarker.
func (GoTool) Name() string {
	return "go tool"
}

// Benchmark runs go benchmarks. Implements Benchmarker.
func (GoTool) Benchmark(workspace *Workspace, regexp string) (string, error) {
	honourVendorExperiment(workspace)

	return Run(workspace.Path(), "go", "test", "-run", "XXXXXX", "-bench", regexp, "-benchmem")
}

// Check guesses if it's normal go project. Basically it checks if
// there are any *.go files in directory. Implements Benchmarker.
func (GoTool) Check(workspace *Workspace) bool {
	files, err := ioutil.ReadDir(workspace.Path())
	if err != nil {
		return false
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".go" {
			return true
		}
	}

	return false
}

// GbTool is a wrapper for 'gb test' tool.
type GbTool struct{}

// Name returns tool name. Implements Benchmarker.
func (GbTool) Name() string {
	return "gb tool"
}

// Benchmark runs gb benchmarks. Implements Benchmarker.
func (GbTool) Benchmark(workspace *Workspace, regexp string) (string, error) {
	honourVendorExperiment(workspace)

	return Run(workspace.Path(), "gb", "test", "-run", "XXXXXX", "-bench", regexp, "-benchmem")
}

// Check guesses if it's normal go progject. Implements Benchmarker.
func (GbTool) Check(workspace *Workspace) bool {
	files, err := ioutil.ReadDir(workspace.Path())
	if err != nil {
		return false
	}

	for _, f := range files {
		if f.Name() == "src" {
			return true
		}
	}

	return false
}
