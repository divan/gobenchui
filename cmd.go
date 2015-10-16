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
}

// Run launches command in the given dir and handles success/errors.
func Run(dir, command string, args ...string) ([]string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, &RunError{
			Message: err.Error(),
			Stderr:  stderr.String(),
		}
	}

	return strings.Split(stdout.String(), "\n"), nil
}

// Error implements error interface for RunError.
func (r *RunError) Error() string {
	if r.Stderr != "" {
		return fmt.Sprintf("git failed: %s", r.Stderr)
	} else {
		return fmt.Sprintf("git failed: %s", r.Message)
	}
}
