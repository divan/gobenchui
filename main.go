package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
)

// ProgramName specifies default program name
// (for tempfiles, etc)
var ProgramName = "gobenchui"

func main() {
	bind := flag.String("bind", ":6222", "host:port to bind http server to")
	benchOpts := flag.String("benchOpts", "", "Custom 'go test' flags")
	flag.Parse()
	if len(flag.Args()) != 1 {
		Usage()
		os.Exit(1)
	}

	pkg := flag.Arg(0)
	path, err := getAbsPath(pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to find package:", err)
		os.Exit(1)
	}
	fmt.Println("Benchmarking package", path)

	// only git so far
	var vcs VCS
	vcs, err = NewGitVCS(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "package isn't under any supported VCS, so no benchmarks to compare\n")
		os.Exit(1)
	}

	err = vcs.Workspace().Clone()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't clone dir:", err)
		os.Exit(1)
	}
	clonedPath := vcs.Workspace().Path()
	fmt.Println("[DEBUG] Cloned package to", clonedPath)

	// Remove temporary directory in the end
	cleanup := func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Couldn't delete temp dir:", err)
		}
	}
	defer cleanup(clonedPath)

	// Prepare commits to run benchmarks agains
	commits, err := vcs.Commits()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get commits:", err)
		return
	}

	ch, err := RunBenchmarks(vcs, commits)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't run benchmarks:", err)
		return
	}

	info := NewInfo(pkg, path, vcs.Name(), *benchOpts, commits)

	go StartServer(*bind, ch, info)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	<-sigCh
	fmt.Println("Got signal, cleaning up")
}

// Usage prints program usage text.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s package\n", os.Args[0])
	flag.PrintDefaults()
}

// getAbsPath returns absolute path to package to be benchmarked.
// For package names it looks for them in GOPATH.
// For '.' it resolves current working directory.
func getAbsPath(pkg string) (string, error) {
	if pkg == "." {
		return os.Getwd()
	}

	path := filepath.Join(GOPATH(), "src", pkg)
	return path, nil
}
