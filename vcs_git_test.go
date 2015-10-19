package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
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

func TestGitParseCommits(t *testing.T) {
	Convey("Parsing commits should work correctly", t, func() {
		Convey("Normal case", func() {
			lines := []string{
				`378739736dd1baa076675c02fe45822bf8936a14|Sun, 18 Oct 2015 22:49:00 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Return again`,
				`4c6378260658f763d67dea3a931a25daf2a82d51|Sun, 18 Oct 2015 22:48:14 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Returned panic`,
				`69814659db92a97e077196207815ca12cba7e0dd|Sun, 18 Oct 2015 22:47:54 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Added no build`,
				`24ebd4283248f85746c69d1d97c7dddedd18d863|Sun, 18 Oct 2015 22:30:15 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Remove panic`,
				`aec0795b436742b5669fdebd28df702704b8afb2|Sun, 18 Oct 2015 22:29:59 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Add panic`,
			}
			commits := parseGitCommits(lines, time.UTC)
			So(len(commits), ShouldEqual, len(lines))
			So(commits[0], ShouldResemble, Commit{
				Hash:    "378739736dd1baa076675c02fe45822bf8936a14",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "Return again",
				Date:    time.Date(2015, time.October, 18, 22, 49, 00, 00, time.UTC),
			})
			So(commits[4], ShouldResemble, Commit{
				Hash:    "aec0795b436742b5669fdebd28df702704b8afb2",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "Add panic",
				Date:    time.Date(2015, time.October, 18, 22, 29, 59, 00, time.UTC),
			})
		})
		Convey("Day < 10 case", func() {
			lines := []string{
				`378739736dd1baa076675c02fe45822bf8936a14|Sun, 8 Oct 2015 22:49:00 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Return again`,
			}
			commits := parseGitCommits(lines, time.UTC)
			So(len(commits), ShouldEqual, len(lines))
			So(commits[0], ShouldResemble, Commit{
				Hash:    "378739736dd1baa076675c02fe45822bf8936a14",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "Return again",
				Date:    time.Date(2015, time.October, 8, 22, 49, 00, 00, time.UTC),
			})
		})
		Convey("Separators in subject case", func() {
			lines := []string{
				`378739736dd1baa076675c02fe45822bf8936a14|Sun, 8 Oct 2015 22:49:00 +0000|Ivan Daniluk <ivan.daniluk@gmail.com>|Return with character'|' and again |[]|`,
			}
			commits := parseGitCommits(lines, time.UTC)
			So(len(commits), ShouldEqual, len(lines))
			So(commits[0], ShouldResemble, Commit{
				Hash:    "378739736dd1baa076675c02fe45822bf8936a14",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "Return with character'|' and again |[]|",
				Date:    time.Date(2015, time.October, 8, 22, 49, 00, 00, time.UTC),
			})
		})
	})
}
