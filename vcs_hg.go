package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Hg implements VCS for Mercurial version control.
type Hg struct {
	workspace *Workspace

	filter FilterOptions
}

// NewHgVCS returns new Mercurial vcs, and checks it it's valid hg workspace.
func NewHgVCS(path string, filter FilterOptions) (*Hg, error) {
	// check if this path is under hg control
	_, err := Run(path, "hg", "verify")
	if err != nil {
		return nil, err
	}

	// find top-level (root) directory of workspace
	out, err := Run(path, "hg", "root")
	if err != nil {
		return nil, errors.New("cannot determine root folder")
	}
	root := strings.TrimSpace(out)
	prefix := findPrefix(path, root)

	workspace := NewWorkspace(root, prefix)
	vcs := &Hg{
		workspace: workspace,
		filter:    filter,
	}
	return vcs, nil
}

// Commits returns all commits for the current branch. Implements VCS interface.
func (g *Hg) Commits() ([]Commit, error) {
	path := g.Workspace().Path()

	// Prepare args, and add user defined args to `hg log` command
	args := []string{"log", `--template={node}%{date|rfc822date}%{author}%{desc}\n`}
	if len(g.filter.Args) > 0 {
		// Append custom arguments, excluding formatting-related ones
		cleanedArgs := cleanHgArgs(g.filter.Args...)
		args = append(args, cleanedArgs...)
	}

	if g.filter.LastN > 0 {
		lastNArg := fmt.Sprintf("-l %d", g.filter.LastN)
		args = append(args, lastNArg)
	}

	out, err := Run(path, "hg", args...)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")

	commits := parseHgCommits(lines, time.Local)

	// Filter to max entries, if specified
	if g.filter.Max > 0 {
		commits = FilterMax(commits, g.filter.Max)
	}

	return commits, nil
}

// parseHgCommits parses output from `hg log` command.
func parseHgCommits(lines []string, location *time.Location) []Commit {
	var commits []Commit
	for _, str := range lines {
		fields := strings.SplitN(str, "%", 4)
		if len(fields) != 4 {
			fmt.Fprintln(os.Stderr, "[ERROR] Wrong commit info, skipping:", len(fields), str)
			continue
		}
		timestamp, err := time.ParseInLocation(time.RFC1123Z, fields[1], location)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Cannot parse timestamp:", err)
			continue
		}
		commit := Commit{
			Hash:    fields[0],
			Date:    timestamp,
			Subject: fields[3],
			Author:  fields[2],
		}
		commits = append(commits, commit)
	}

	return commits
}

// SwitchTo switches to the given commit by hash. Implements VCS interface.
func (g *Hg) SwitchTo(hash string) error {
	path := g.Workspace().Path()
	_, err := Run(path, "hg", "update", hash)
	return err
}

// Workspace returns assosiated Workspace. Implements VCS interface.
func (g *Hg) Workspace() *Workspace {
	return g.workspace
}

// Name returns vcs common name. Implements VCS interface.
func (*Hg) Name() string {
	return "hg"
}

var (
	// TODO: find a person who use mercurial a lot and can
	// help to add more ignored args (ones that may affect
	// predetermined output)
	ignoredHgArgs = []string{
		"--template",
	}
)

// cleanHgArgs cleans user defined custom hg arguments.
// it basically removes arguments, that may affect formatting
// output (we use specific format for parsing results)
func cleanHgArgs(args ...string) []string {
	var ret []string
	for _, arg := range args {
		trimmed := strings.TrimSpace(arg)
		if trimmed == "" {
			continue
		}

		var ignore bool
		for _, ignored := range ignoredHgArgs {
			if strings.HasPrefix(trimmed, ignored) {
				ignore = true
			}
		}
		if !ignore {
			ret = append(ret, trimmed)
		}
	}
	return ret
}
