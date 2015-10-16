package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Git implements VCS for Git version control.
type Git struct {
	path string
}

// Commits returns all commits for the current branch. Implements VCS interface.
func (g *Git) Commits() ([]Commit, error) {
	out, err := Run(g.Path(), "git", "log", `--pretty=format:"%H|%cd|%s|%cn %ce"`, `--date=rfc`)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")

	var commits []Commit
	for _, str := range lines {
		fields := strings.Split(str, "|")
		if len(fields) != 4 {
			fmt.Fprintln(os.Stderr, "Wrong commit info, skipping: %s", str)
			continue
		}
		timestamp, err := time.Parse(time.RFC1123Z, fields[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot parse timestamp: %v", err)
			continue
		}
		commit := Commit{
			Hash:    fields[0],
			Date:    timestamp,
			Subject: fields[2],
			Author:  fields[3],
		}
		commits = append(commits, commit)
	}

	return commits, nil
}

func (g *Git) SwitchTo(hash string) error {
	return nil
}
func (g *Git) PreviousCommit() string {
	return ""
}
func (g *Git) Path() string {
	return g.path
}
func (g *Git) SetPath(path string) {
	g.path = path
}
