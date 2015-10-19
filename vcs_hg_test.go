package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestHgParseCommits(t *testing.T) {
	Convey("Parsing commits should work correctly", t, func() {
		Convey("Normal case", func() {
			lines := []string{
				`05a2ab6a17d3ef8ea71cddb8e4c56ce1cfcd53aa%Mon, 19 Oct 2015 02:33:30 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%new test`,
				`9ef3f9bab2683ff1998ec129e4b3b3a26b95ca9b%Mon, 19 Oct 2015 02:13:25 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%add bench`,
				`1cda6ab6b89db97aeeb27f96c61abff211956471%Thu, 14 May 2015 14:19:58 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%Updated README again`,
				`557ab4832839da1195a5e82c168bd5899570f37f%Thu, 14 May 2015 14:13:41 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%README.md edited online with Bitbucket`,
				`590df12383190d7984942de8f5503076da66d3b2%Thu, 14 May 2015 16:20:28 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%Added fmt.`,
				`47f015f5e63bd9018b7d581b98594f8126ca5609%Thu, 14 May 2015 16:12:55 +0000%Ivan Daniluk <ivan.daniluk@gmail.com>%Initial commit`,
			}
			commits := parseHgCommits(lines, time.UTC)
			So(len(commits), ShouldEqual, len(lines))
			So(commits[0], ShouldResemble, Commit{
				Hash:    "05a2ab6a17d3ef8ea71cddb8e4c56ce1cfcd53aa",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "new test",
				Date:    time.Date(2015, time.October, 19, 02, 33, 30, 00, time.UTC),
			})
			So(commits[len(commits)-1], ShouldResemble, Commit{
				Hash:    "47f015f5e63bd9018b7d581b98594f8126ca5609",
				Author:  "Ivan Daniluk <ivan.daniluk@gmail.com>",
				Subject: "Initial commit",
				Date:    time.Date(2015, time.May, 14, 16, 12, 55, 00, time.UTC),
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
