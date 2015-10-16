package main

import (
	"fmt"
	"os"
)

type Benchmark string

func RunBenchmarks(vcs VCS) (chan Benchmark, error) {
	ch := make(chan Benchmark)

	commits, err := vcs.Commits()
	if err != nil {
		return nil, err
	}

	go func(commits []Commit) {
		path := vcs.Workspace().Path()
		for _, commit := range commits {
			// Switch to previous commit
			fmt.Printf("Switching to |%s|\n", commit.Hash)
			_, err := Run(path, "git", "checkout", commit.Hash)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			// Run benchmark for this commit
			out, err := Run(path, "go", "test", "-test.bench", ".")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			ch <- Benchmark(out)
		}
	}(commits)

	return ch, nil
}
