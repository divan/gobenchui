package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGitCleanOptions(t *testing.T) {
	Convey("CleanGitArgs should clean git args correctly", t, func() {
		Convey("clean up --pretty args", func() {
			args := []string{"--pretty=oneline"}
			cleaned := cleanGitArgs(args...)
			So(len(cleaned), ShouldEqual, 0)
		})
		Convey("clean up --date args", func() {
			args := []string{"--date=iso"}
			cleaned := cleanGitArgs(args...)
			So(len(cleaned), ShouldEqual, 0)
		})
		Convey("do not clean up harmless flags", func() {
			args := []string{"--date=iso", "--author=Ivan"}
			cleaned := cleanGitArgs(args...)
			So(len(cleaned), ShouldEqual, 1)
			So(cleaned[0], ShouldEqual, "--author=Ivan")
		})
	})
}
