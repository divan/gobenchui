package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Workspace represents local top-level VCS workspace.
//
// Package may be in subfolder (like github.com/etcd/coreos/store),
// but Workspace represents the whole directory (github.com/etcd/coreos).
//
// Following this example, root is '~/github.com/etcd/coreos',
// prefix is 'store'.
//
// gopath is introduced to handle GO15VENDOREXPERIMENT gopath issue,
// if root is '/tmp/tempXXXX/src/pkg/github.com/etcd/coreos',
// then gopath is '/tmp/tempXXXX'. Optional.
type Workspace struct {
	root   string
	prefix string

	gopath string
}

// NewWorkspace creates new Workspace.
func NewWorkspace(root, prefix string) *Workspace {
	return &Workspace{
		root:   root,
		prefix: prefix,
	}
}

// Path returns full workspace path to package (w/ prefix).
func (w *Workspace) Path() string {
	return filepath.Join(w.root, w.prefix)
}

// Root returns root directory for workspace (w/o prefix).
func (w *Workspace) Root() string {
	return w.root
}

// Gopath returns root gopath directory for workspace (w/o prefix).
func (w *Workspace) Gopath() string {
	return w.gopath
}

// SetRoot sets new root path for workspace.
func (w *Workspace) SetRoot(gopath, root string) {
	w.gopath = gopath
	w.root = root
}

// Clone copies whole workspace to temporary directory.
func (w *Workspace) Clone() error {
	tmp, err := ioutil.TempDir("", ProgramName)
	if err != nil {
		return err
	}

	fmt.Println("[DEBUG] Cloning git workspace to", tmp)
	// place sources under src/pkg to make it look like
	// proper GOPATH (needed for GO15VENDOREXPERIMENT support)
	targetDir := filepath.Join(tmp, "src", "pkg")
	err = os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		os.RemoveAll(tmp)
		return err
	}
	err = copyAll(targetDir+"/", w.Root())
	if err != nil {
		os.RemoveAll(tmp)
		return err
	}
	w.SetRoot(tmp, targetDir)

	return nil
}

// copyFile copies the file with path src to dst. The new file must not exist.
// It is created with the same permissions as src.
func copyFile(dst, src string) error {
	rf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer rf.Close()
	rstat, err := rf.Stat()
	if err != nil {
		return err
	}
	if rstat.IsDir() {
		return fmt.Errorf("dir argument to CopyFile (%s, %s)", dst, src)
	}

	wf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, rstat.Mode())
	if err != nil {
		return err
	}
	if _, err := io.Copy(wf, rf); err != nil {
		wf.Close()
		return err
	}
	return wf.Close()
}

// copyAll copies the file or (recursively) the directory at src to dst.
// Permissions are preserved. dst must already exist.
func copyAll(dst, src string) error {
	return filepath.Walk(src, makeWalkFn(dst, src))
}

func makeWalkFn(dst, src string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dir := strings.TrimPrefix(path, src)
		if dir == "/" || dir == "" {
			return nil
		}
		dstPath := filepath.Join(dst, dir)
		if info.IsDir() {
			return os.Mkdir(dstPath, info.Mode())
		}
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if newPath, err := os.Readlink(path); err != nil {
				return err
			} else {
				if newPath[0] != '/' {
					// Relative symlink
					path = filepath.Join(filepath.Dir(path), newPath)
				} else {
					path = newPath
				}

				if info, err = os.Lstat(path); err != nil {
					return err
				}

				if info.IsDir() {
					if err = os.Mkdir(dstPath, info.Mode()); err != nil {
						return err
					}
					// Following the dir symlink
					return filepath.Walk(path, makeWalkFn(dstPath, path))
				}
			}
		}
		return copyFile(dstPath, path)
	}
}
