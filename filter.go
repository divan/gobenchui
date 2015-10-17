package main

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
