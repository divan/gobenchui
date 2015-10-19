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

	// Remove temporary directory in the end
	cleanup := func() {
		path := vcs.Workspace().Root()
		os.RemoveAll(path)
	}
	defer cleanup()

	// Prepare commits to run benchmarks agains
	commits, err := vcs.Commits()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get commits:", err)
		return
	}

	ch := RunBenchmarks(vcs, commits, *benchArgs)

	info := NewInfo(pkg, path, vcs.Name(), *benchArgs, commits)
	info.SetStatus(InProgress)

	// There is basically no reason to make this channel
	// buffered, but just in case, if web frontend code will
	// stuck (websocket js issue or smth.), results will
	// still be saved into info, so the page reload will
	// show all results.
	webCh := make(chan interface{}, 1024)
	go func() {
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					info.SetStatus(Finished)
					info.SetCommit(nil)
					webCh <- BenchmarkStatus{
						Status:   Finished,
						Progress: 100.0,
					}
					fmt.Println("[INFO] Finished, cleaning up.")
					cleanup()
					return
				}
				if result, ok := val.(BenchmarkSet); ok {
					info.AddResult(result)
					info.SetStatus(InProgress)
				}
				if status, ok := val.(BenchmarkRun); ok {
					info.SetStatus(InProgress)
					info.SetCommit(&status.Commit)
				}

				webCh <- val
			}
		}
	}()

	go StartServer(*bind, webCh, info)

	// don't exit, even after all benchmarks had been completed,
	// as we need to keep serve web page
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	<-sigCh
	fmt.Println("Got signal, exiting...")
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
