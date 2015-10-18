package main

import (
	"sync"
	"time"
)

type Status string

const (
	Undefined  Status = "Undefined"
	InProgress        = "In progress"
	Finished          = "Finished"
	Failed            = "Failed"
)

// Info holds information about bench session,
// like pkg name, start time, progress, status, etc.
type Info struct {
	mx *sync.RWMutex

	Status   Status  `json:"status"`
	Progress float64 `json:"progress"`

	PkgName string `json:"pkg_name"`
	PkgPath string `json:"pkg_path"`
	VCS     string `json:"vcs"`

	BenchOptions string   `json:"bench_options"`
	Commits      []Commit `json:"commits"`

	BenchResults []BenchmarkSet  `json:"results"`
	Series       *HighchartsData `json:"series,omitempty"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// NewInfo returns new initialized info.
func NewInfo(pkg, path, vcs, benchopts string, commits []Commit) *Info {
	return &Info{
		mx: &sync.RWMutex{},

		Status:   Undefined,
		Progress: 0.0,

		PkgName: pkg,
		PkgPath: path,
		VCS:     vcs,

		BenchOptions: benchopts,
		Commits:      commits,

		StartTime: time.Now(),
	}
}

// SetProgress is a setter for Progress value.
func (i *Info) SetProgress(v float64) {
	i.mx.Lock()
	i.Progress = v
	i.mx.Unlock()
}

// SetStatus changes status of execution.
func (i *Info) SetStatus(status Status) {
	i.mx.Lock()
	defer i.mx.Unlock()
	i.Status = status
	if status == Finished {
		i.EndTime = time.Now()
	}
}

// AddResults inserts new BenchmarkSet result to Info.
func (i *Info) AddResult(b BenchmarkSet) {
	i.mx.Lock()
	defer i.mx.Unlock()
	i.BenchResults = append(i.BenchResults, b)
	if i.Series == nil {
		i.Series = &HighchartsData{}
	}
	i.Series.AddResult(b)
}
