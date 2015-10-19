package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFilterMax(t *testing.T) {
	Convey("Filter Max should filter commits evenly", t, func() {
		N := int64(100)
		var commits []Commit
		for i := int64(0); i < N; i++ {
			date, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("2015-%02d-%02d 18:%02d:01", i%12+1, i%28+1, i%60))
			commit := Commit{
				Hash:    fmt.Sprintf("Hash %d", i),
				Author:  fmt.Sprintf("Author %d", i%3),
				Subject: fmt.Sprintf("Subject %d", i),
				Date:    date,
			}
			commits = append(commits, commit)
		}
		Convey("Special case, len == 0", func() {
			res := FilterMax([]Commit{}, 0)
			So(len(res), ShouldEqual, 0)
		})
		Convey("Special case, max > len", func() {
			res := FilterMax(commits, N*2)
			So(len(res), ShouldEqual, N)
		})
		Convey("Special case, max == 0", func() {
			res := FilterMax(commits, 0)
			So(len(res), ShouldEqual, N)
		})
		Convey("Max = 2, start and end", func() {
			res := FilterMax(commits, 2)
			So(len(res), ShouldEqual, 2)
			So(res[0].Hash, ShouldEqual, "Hash 0")
			So(res[1].Hash, ShouldEqual, "Hash 99")
		})
		Convey("Max = 3, start, middle and end", func() {
			res := FilterMax(commits, 3)
			So(len(res), ShouldEqual, 3)
			So(res[0].Hash, ShouldEqual, "Hash 0")
			So(res[1].Hash, ShouldEqual, "Hash 49")
			So(res[2].Hash, ShouldEqual, "Hash 99")
		})
		Convey("Max = 10, tens", func() {
			res := FilterMax(commits, 10)
			So(len(res), ShouldEqual, 10)
			So(res[0].Hash, ShouldEqual, "Hash 0")
			So(res[1].Hash, ShouldEqual, "Hash 11")
			So(res[2].Hash, ShouldEqual, "Hash 22")
			So(res[3].Hash, ShouldEqual, "Hash 33")
			So(res[4].Hash, ShouldEqual, "Hash 44")
			So(res[5].Hash, ShouldEqual, "Hash 55")
			So(res[6].Hash, ShouldEqual, "Hash 66")
			So(res[7].Hash, ShouldEqual, "Hash 77")
			So(res[8].Hash, ShouldEqual, "Hash 88")
			So(res[9].Hash, ShouldEqual, "Hash 99")
		})
		Convey("Max = 98, near len", func() {
			res := FilterMax(commits, 98)
			So(len(res), ShouldEqual, 98)
			So(res[0].Hash, ShouldEqual, "Hash 0")
			So(res[90].Hash, ShouldEqual, "Hash 90")
			So(res[96].Hash, ShouldEqual, "Hash 97")
			So(res[97].Hash, ShouldEqual, "Hash 99")
		})
		Convey("Max = 75, 3/4", func() {
			res := FilterMax(commits, 75)
			So(len(res), ShouldEqual, 75)
			So(res[0].Hash, ShouldEqual, "Hash 0")
			So(res[1].Hash, ShouldEqual, "Hash 1")
			So(res[73].Hash, ShouldEqual, "Hash 97")
			So(res[74].Hash, ShouldEqual, "Hash 99")
		})
	})
	Convey("Filter string() should generate text correctly", t, func() {
		Convey("When nothing specified", func() {
			str := NewFilterOptions(0, 0).String()
			So(str, ShouldEqual, "all commits")
		})
		Convey("When lastN and Max specified", func() {
			str := NewFilterOptions(99, 20).String()
			So(str, ShouldEqual, "max 20 from last 99 commits")
		})
		Convey("When lastN specified", func() {
			str := NewFilterOptions(150, 0).String()
			So(str, ShouldEqual, "last 150 commits")
		})
		Convey("When Max specified", func() {
			str := NewFilterOptions(0, 25).String()
			So(str, ShouldEqual, "max 25 from all commits")
		})
	})
}
