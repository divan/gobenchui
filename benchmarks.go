package main

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"os"
	"time"
)

// BenchmarkSet represents a set of benchmarks for single commit.
type BenchmarkSet struct {
	Commit Commit    `json:"commit"`
	Set    parse.Set `json:"set"`
	Error  error     `json:"-"`
}

// BenchmarkRun represents current state of benchmark being run.
type BenchmarkRun struct {
	Commit    Commit    `json:"commit,omitempty"`
	Error     error     `json:"error,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

// RunBenchmarks loops over given commits and runs benchmarks for each of them.
func RunBenchmarks(vcs VCS, commits []Commit, benchRegexp string) chan interface{} {
	ch := make(chan interface{})

	go func(commits []Commit) {
		defer close(ch)

		handleError := func(err error, run BenchmarkRun) {
			fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
			run.Error = err
			ch <- run
		}

		for _, commit := range commits {
			run := BenchmarkRun{
				Commit:    commit,
				StartTime: time.Now(),
			}
			ch <- run

			// Switch to previous commit
			fmt.Printf("[DEBUG] Switching to %s\n", commit.Hash)
			if err := vcs.SwitchTo(commit.Hash); err != nil {
				handleError(err, run)
				return
			}

			// Run benchmark for this commit
			// but first, try to guess what vendoring is used (if any)
			// and use appropriate tool
			runBenchmark := func(b Benchmarker) (string, error, bool) {
				if !b.Check(vcs.Workspace()) {
					return "", nil, false
				}
				fmt.Println("[DEBUG] Using", b.Name())
				out, err := b.Benchmark(vcs.Workspace(), benchRegexp)
				return out, err, true
			}
			// Try normal go tool
			out, err, ok := runBenchmark(GoTool{})
			if !ok {
				// Try GB
				out, err, _ = runBenchmark(GbTool{})
			}
			if err != nil {
				handleError(err, run)
				continue
			}

			set, err := ParseBenchmarkOutput(out)
			if err != nil {
				handleError(err, run)
				continue
			}

			set.Commit = commit

			ch <- *set
		}
	}(commits)

	return ch
}

// ParseBenchmarkOutput parses raw output from 'go test -test.bench' command.
func ParseBenchmarkOutput(out string) (*BenchmarkSet, error) {
	buf := bytes.NewBufferString(out)
	set, err := parse.ParseSet(buf)
	if err != nil {
		return nil, err
	}
	return &BenchmarkSet{
		Set: set,
	}, nil
}
