package main

import (
	"time"
)

// VCS represents our needs from VCS - list commits, switch
// to specific commit and obtain previous commit.
type VCS interface {
	Commits() ([]Commit, error)
	SwitchTo(hash string) error
	Workspace() *Workspace
}

// Commit represents single commit in VCS.
//
// Author is the last person touched this commit
// (committer in git terms)
type Commit struct {
	Hash    string    `json:"hash"`
	Author  string    `json:"author"`
	Subject string    `json:"subject"`
	Date    time.Time `json:"date"`
}
