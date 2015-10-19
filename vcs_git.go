package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Git implements VCS for Git version control.
type Git struct {
	workspace *Workspace

	filter FilterOptions
}

// NewGitVCS returns new Git vcs, and checks it it's valid git workspace.
func NewGitVCS(path string, filter FilterOptions) (*Git, error) {
	// check if this path is under git control
	_, err := Run(path, "git", "rev-parse", "--git-dir")
	if err != nil {
		return nil, err
	}

	// find top-level (root) directory of workspace
	out, err := Run(path, "git", "rev-parse", "--show-toplevel")
	if err != nil {
		return nil, errors.New("cannot determine root folder")
	}
	root := strings.TrimSpace(out)
	prefix := findPrefix(path, root)

	workspace := NewWorkspace(root, prefix)
	vcs := &Git{
		workspace: workspace,
		filter:    filter,
	}
	return vcs, nil
}

// Commits returns all commits for the current branch. Implements VCS interface.
func (g *Git) Commits() ([]Commit, error) {
	path := g.Workspace().Path()

	// Prepare args, and add user defined args to `git log` command
	args := []string{"log", `--pretty=format:%H|%cd|%cn <%ce>|%s`, `--date=rfc`}
	if len(g.filter.Args) > 0 {
		// Append custom arguments, excluding formatting-related ones
		cleanedArgs := cleanGitArgs(g.filter.Args...)
		args = append(args, cleanedArgs...)
	}

	if g.filter.LastN > 0 {
		lastNArg := fmt.Sprintf("-n %d", g.filter.LastN)
		args = append(args, lastNArg)
	}

	out, err := Run(path, "git", args...)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")
	commits := parseGitCommits(lines, time.Local)

	// Filter to max entries, if specified
	if g.filter.Max > 0 {
		commits = FilterMax(commits, g.filter.Max)
	}

	return commits, nil
}

// parseGitCommits parses output from `git log` command.
func parseGitCommits(lines []string, location *time.Location) []Commit {
	var commits []Commit
	for _, str := range lines {
		fields := strings.SplitN(str, "|", 4)
		if len(fields) != 4 {
			fmt.Fprintln(os.Stderr, "[ERROR] Wrong commit info, skipping:", len(fields), str)
			continue
		}
		timestamp, err := time.ParseInLocation(RFC1123ZGit, fields[1], location)
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
func (g *Git) SwitchTo(hash string) error {
	path := g.Workspace().Path()
	_, err := Run(path, "git", "checkout", hash)
	return err
}

// Workspace returns assosiated Workspace. Implements VCS interface.
func (g *Git) Workspace() *Workspace {
	return g.workspace
}

// Name returns vcs common name. Implements VCS interface.
func (*Git) Name() string {
	return "git"
}

var (
	ignoredGitArgs = []string{
		"--date-order", "--author-date-order", "--topo-order", "--reverse",
		"--relative-date", "--date", "--parents", "--children", "--left-right",
		"--graph", "--show-linear-break", "--pretty",
	}
)

// cleanGitArgs cleans user defined custom git arguments.
// it basically removes arguments, that may affect formatting
// output (we use specific format for parsing results)
func cleanGitArgs(args ...string) []string {
	var ret []string
	for _, arg := range args {
		trimmed := strings.TrimSpace(arg)
		if trimmed == "" {
			continue
		}

		var ignore bool
		for _, ignored := range ignoredGitArgs {
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

// RFC1123ZGit is a git variation of RFC1123 time layout (--date=rfc)
const RFC1123ZGit = "Mon, 2 Jan 2006 15:04:05 -0700"
