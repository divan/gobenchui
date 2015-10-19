package main

import (
	"fmt"
)

// FilterOptions represents advanced filtering for vcs commits.
type FilterOptions struct {
	// LastN is for 'last N commits'
	// Default: 0 (disabled)
	LastN int64
	// Max defines how many maximum commits should be taken.
	// If total commits number is bigger than Max,
	// it tries to pick commits evenly, so benchmark result
	// can be representable in overall.
	// Default: 100 (disabled)
	Max int64

	Args []string
}

// NewFilterOptions creates new FilterOptions.
func NewFilterOptions(lastN, max int64, args ...string) *FilterOptions {
	return &FilterOptions{
		LastN: lastN,
		Max:   max,
		Args:  args,
	}
}

// FilterMax filters commits, selecting at most 'max'
// evenly distributed items.
//
// Idea is to get brief overview of benchmark progress
// overtime for codebases with large number of commits.
// If there are 10K commits and max=100, it will pick
// 100 commits from very beginning to the last one, with
// the equal time intervals between each picked commit.
func FilterMax(commits []Commit, max int64) []Commit {
	// special cases
	if max < 2 {
		return commits
	}
	length := int64(len(commits))
	if max >= length || length == 0 {
		return commits
	}

	// rough naive implementation (by commits number)
	var ret []Commit
	for i := int64(0); i < length; {
		// how many items left in target slice
		left := max - int64(len(ret)) - 1
		if left == 0 {
			left = 1
		}
		// size of next step
		size := (length - i - 1) / left
		if size == 0 {
			size = 1
		}

		ret = append(ret, commits[i])
		i = i + size
	}
	return ret
}

// String implements Stringer for FilterOptions.
func (f *FilterOptions) String() string {
	out := ""
	if f.Max > 0 {
		out = fmt.Sprintf("%smax %d from ", out, f.Max)
	}
	if f.LastN > 0 {
		out = fmt.Sprintf("%slast %d ", out, f.LastN)
	} else {
		out = fmt.Sprintf("%sall ", out)
	}
	out = fmt.Sprintf("%scommits", out)
	return out
}
