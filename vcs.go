package main

import (
	"os"
	"path/filepath"
	"time"
)

// VCS represents our needs from VCS - list commits, switch
// to specific commit and obtain previous commit.
type VCS interface {
	Commits() ([]Commit, error)
	SwitchTo(hash string) error
	PreviousCommit() string
}

// Commit represents single commit in VCS.
//
// Author is the last person touched this commit
// (committer in git terms)
type Commit struct {
	Hash    string
	Author  string
	Subject string
	Date    time.Time
}

// GetVCS checks whether given path is under VCS control
// and returns appropriate VCS interface implementation.
// TODO: make more sophisticated checks
func GetVCS(path string) (VCS, error) {
	// try git
	gitpath := filepath.Join(path, ".git")
	_, err := os.Stat(gitpath)
	if err != nil {
		return nil, err
	}
	vcs := &Git{
		Path: path,
	}
	return vcs, err
}
