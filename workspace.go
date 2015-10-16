package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Workspace represents local VCS workspace.
type Workspace struct {
	path string
}

// NewWorkspace creates new Workspace.
func NewWorkspace(path string) *Workspace {
	return &Workspace{
		path: path,
	}
}

// Path returns local workspace path.
func (w *Workspace) Path() string {
	return w.path
}

// SetPath sets new path for workspace.
func (w *Workspace) SetPath(path string) {
	w.path = path
}

// Clone copies whole workspace to temporary directory.
func (w *Workspace) Clone() error {
	tmp, err := ioutil.TempDir("", ProgramName)
	if err != nil {
		return err
	}

	err = copyAll(tmp+"/", w.Path())
	if err != nil {
		return err
	}
	w.SetPath(tmp)

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
		return errors.New("dir argument to CopyFile")
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
		return copyFile(dstPath, path)
	}
}
