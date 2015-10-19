package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// RunError represents command running error
type RunError struct {
	Message string
	Stderr  string
	Type    RunErrorType
}

// RunErrorType is a subtype for run error.
type RunErrorType string

const (
	PanicErr        RunErrorType = "panic"
	BuildFailedErr               = "build_failed"
	NoBenchmarksErr              = "no_benchmarks"
	OtherErr                     = "other"
)

// Run launches command in the given dir and handles success/errors.
func Run(dir, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", &RunError{
			Type:    guessErrType(err, stderr.String()),
			Message: err.Error(),
			Stderr:  stderr.String(),
		}
	}

	return stdout.String(), nil
}

// guessErrType tries to guess error type based on stderr and other info.
func guessErrType(err error, stderr string) RunErrorType {
	if strings.HasPrefix(stderr, "panic:") {
		return PanicErr
	}
	if strings.HasPrefix(stderr, "# ") || strings.HasPrefix(stderr, "can't load package") {
		return BuildFailedErr
	}
	return OtherErr
}

// Error implements error interface for RunError.
func (r *RunError) Error() string {
	if r.Stderr != "" {
		return fmt.Sprintf("failed: %s", r.Stderr)
	}

	return fmt.Sprintf("failed: %s", r.Message)
}
