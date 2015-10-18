package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
)

var (
	// ProgramName specifies default program name
	// (for tempfiles, etc)
	ProgramName = "gobenchui"

	bind      = flag.String("bind", ":6222", "host:port to bind http server to")
	vcsArgs   = flag.String("vcsArgs", "", "Additional args for vcs command (git, hg, etc)")
	benchArgs = flag.String("bench", ".", "Regexp for benchmarks, as for `go test -bench`")
	lastN     = flag.Int64("n", 0, "Last N commits only")
	max       = flag.Int64("max", 0, "Maximum commits (distribute evenly)")
)

func main() {
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

	var vcs VCS
	filter := NewFilterOptions(*lastN, *max, *vcsArgs)
	// only git so far
	vcs, err = NewGitVCS(path, *filter)
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
	defer cleanup(vcs.Workspace().Root())

	// Prepare commits to run benchmarks agains
	commits, err := vcs.Commits()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get commits:", err)
		return
	}

	resultCh, runCh := RunBenchmarks(vcs, commits, *benchArgs)

	info := NewInfo(pkg, path, vcs.Name(), *benchArgs, commits)
	info.SetStatus(InProgress)

	// There is basically no reason to make this channel
	// buffered, but just in case, if web frontend code will
	// stuck (websocket js issue or smth.), results will
	// still be saved into info, so the page reload will
	// show all results.
	webCh := make(chan BenchmarkSet, 256)
	go func() {
		for {
			select {
			case result, ok := <-resultCh:
				if !ok {
					info.SetStatus(Finished)
					return
				}
				info.AddResult(result)
				info.SetStatus(InProgress)

				webCh <- result
			}
		}
	}()

	go StartServer(*bind, webCh, runCh, info)

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
