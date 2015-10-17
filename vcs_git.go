package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Git implements VCS for Git version control.
type Git struct {
	workspace *Workspace
}

// NewGitVCS returns new Git vcs, and checks it it's valid git workspace.
func NewGitVCS(path string) (*Git, error) {
	// check if given path contains .git/ directory
	// TODO: add more sophisiticated workspace check
	gitpath := filepath.Join(path, ".git")
	_, err := os.Stat(gitpath)
	if err != nil {
		return nil, err
	}

	workspace := NewWorkspace(path)
	vcs := &Git{
		workspace: workspace,
	}
	return vcs, nil
}

// Commits returns all commits for the current branch. Implements VCS interface.
func (g *Git) Commits() ([]Commit, error) {
	path := g.Workspace().Path()
	out, err := Run(path, "git", "log", `--pretty=format:%H|%cd|%s|%cn <%ce>`, `--date=rfc`)
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

// SwitchTo switches to the given commit by hash. Implements VCS interface.
func (g *Git) SwitchTo(hash string) error {
	path := g.Workspace().Path()
	_, err := Run(path, "git", "checkout", hash)
	return err
}

// Workspace returns assosiated Workspace. Implements VCS interface.
func (g *Git) Workspace() *Workspace {
	return g.workspace
}
