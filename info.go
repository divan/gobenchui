package main

import (
	"sync"
	"time"
)

type Status string

const (
	Starting   Status = "Starting"
	InProgress        = "In progress"
	Finished          = "Finished"
	Failed            = "Failed"
)

// BenchmarkStatus holds details of current status.
type BenchmarkStatus struct {
	Status        Status  `json:"status"`
	Progress      float64 `json:"progress"`
	CurrentCommit *Commit `json:"commit"`
}

// Info holds information about bench session,
// like pkg name, start time, progress, status, etc.
type Info struct {
	mx *sync.RWMutex

	BenchmarkStatus

	PkgName string `json:"pkg_name"`
	PkgPath string `json:"pkg_path"`
	VCS     string `json:"vcs"`

	BenchOptions string   `json:"bench_options"`
	Commits      []Commit `json:"commits"`

	BenchResults []BenchmarkSet  `json:"results"`
	TimeSeries   *HighchartsData `json:"time_series,omitempty"`
	MemSeries    *HighchartsData `json:"mem_series,omitempty"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// NewInfo returns new initialized info.
func NewInfo(pkg, path, vcs, benchopts string, commits []Commit) *Info {
	return &Info{
		mx: &sync.RWMutex{},

		BenchmarkStatus: BenchmarkStatus{
			Status:   Starting,
			Progress: 0.0,
		},

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

// SetCommit is a setter for Current Commit value.
func (i *Info) SetCommit(commit *Commit) {
	i.mx.Lock()
	i.CurrentCommit = commit
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
	if i.TimeSeries == nil {
		i.TimeSeries = &HighchartsData{}
	}
	i.TimeSeries.AddResult(b, "time")

	if i.MemSeries == nil {
		i.MemSeries = &HighchartsData{}
	}
	i.MemSeries.AddResult(b, "memory")
}
