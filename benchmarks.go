package main

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"os"
)

// BenchmarkSet represents a set of benchmarks for single commit.
type BenchmarkSet struct {
	Commit Commit
	Set    parse.Set
}

// RunBenchmarks loops over given commits and runs benchmarks for each of them.
func RunBenchmarks(vcs VCS, commits []Commit) (chan BenchmarkSet, error) {
	ch := make(chan BenchmarkSet)

	go func(commits []Commit) {
		defer close(ch)

		path := vcs.Workspace().Path()
		for _, commit := range commits {
			// Switch to previous commit
			fmt.Printf("[DEBUG] Switching to %s\n", commit.Hash)
			if err := vcs.SwitchTo(commit.Hash); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			// Run benchmark for this commit
			out, err := Run(path, "go", "test", "-test.bench", ".")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			set, err := ParseBenchmarkOutput(out)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			set.Commit = commit

			ch <- *set
		}
	}(commits)

	return ch, nil
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
