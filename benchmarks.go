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
}

// BenchmarkRun represents current state of benchmark being run.
type BenchmarkRun struct {
	Commit    Commit    `json:"commit,omitempty"`
	Error     error     `json:"error,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

// RunBenchmarks loops over given commits and runs benchmarks for each of them.
func RunBenchmarks(vcs VCS, commits []Commit, benchRegexp string) (chan BenchmarkSet, chan BenchmarkRun) {
	resultCh := make(chan BenchmarkSet)
	runCh := make(chan BenchmarkRun)

	go func(commits []Commit) {
		defer close(resultCh)
		defer close(runCh)

		handleError := func(err error, run BenchmarkRun, ch chan BenchmarkRun) {
			fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
			run.Error = err
			ch <- run
		}

		path := vcs.Workspace().Path()
		for _, commit := range commits {
			run := BenchmarkRun{
				Commit:    commit,
				StartTime: time.Now(),
			}
			runCh <- run

			// Switch to previous commit
			fmt.Printf("[DEBUG] Switching to %s\n", commit.Hash)
			if err := vcs.SwitchTo(commit.Hash); err != nil {
				handleError(err, run, runCh)
				return
			}

			// Run benchmark for this commit
			// TODO: make it command agnostic (for gb and others)
			out, err := Run(path, "go", "test", "-run", "XXXXXX", "-bench", benchRegexp)
			if err != nil {
				handleError(err, run, runCh)
				continue
			}

			set, err := ParseBenchmarkOutput(out)
			if err != nil {
				handleError(err, run, runCh)
				continue
			}

			set.Commit = commit

			resultCh <- *set
		}
	}(commits)

	return resultCh, runCh
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
