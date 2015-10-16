package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// ProgramName specifies default program name
// (for tempfiles, etc)
var ProgramName = "gobenchui"

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		Usage()
		os.Exit(1)
	}

	pkg := flag.Arg(0)
	path, err := getPath(pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to find package:", err)
		os.Exit(1)
	}
	fmt.Println("Benchmarking package", path)

	vcs, err := GetVCS(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "package isn't under any supported VCS, so no benchmarks to compare\n")
		os.Exit(1)
	}

	newPath, err := CloneDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't clone dir:", err)
		os.Exit(1)
	}
	fmt.Println("[DEBUG] Cloned package to", newPath)

	// Replace VCS path with newly copied directory
	vcs.SetPath(newPath)

	// Remove temporary directory in the end
	defer func() {
		err := os.RemoveAll(newPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Couldn't delete temp dir:", err)
		}
	}()

	ch, err := RunBenchmarks(vcs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't run benchmarks:", err)
		os.Exit(1)
	}
	out := <-ch
	fmt.Println("Benchmark results:")
	fmt.Println(out)

}

// Usage prints program usage text.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s package\n", os.Args[0])
	flag.PrintDefaults()
}

// getPath returns absolute path to package to be benchmarked.
// For package names it looks for them in GOPATH.
// For '.' it resolves current working directory.
func getPath(pkg string) (string, error) {
	if pkg == "." {
		return os.Getwd()
	}

	path := filepath.Join(GOPATH(), "src", pkg)
	return path, nil
}
