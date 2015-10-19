package main

import (
	"os"
	"path/filepath"
	"strings"
)

// absPath returns absolute path to package to be benchmarked.
// For package names it looks for them in GOPATH.
// For '.' it resolves current working directory.
func absPath(pkg, gopath string) (string, error) {
	if pkg == "." {
		return os.Getwd()
	}

	path := filepath.Join(gopath, "src", pkg)
	return filepath.Clean(path), nil
}

// normalizePkgName guesses correct package name
// to be shown on UI.
// Usually package is referenced by proper name,
// but when using dot directory, we need to guess it.
func normalizePkgName(pkg, absPath, gopath string) string {
	if pkg != "." {
		return pkg
	}

	prefix := filepath.Join(gopath, "src") + string(filepath.Separator)
	if strings.HasPrefix(absPath, prefix) {
		return strings.TrimPrefix(absPath, prefix)
	}

	// else, it means project is in it's own directory,
	// outside of GOPATH scope, so just use it's basename
	return filepath.Base(absPath)
}

// findPrefix find relative dir in top-level dir.
// It's basically the same as `git rev-parse --show-prefix`.
//
// If package is in top-level directory, prefix is "".
func findPrefix(path, root string) string {
	prefix := strings.TrimPrefix(path, root)
	prefix = strings.TrimPrefix(prefix, "/")
	return prefix
}
