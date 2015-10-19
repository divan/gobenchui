package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPathHelpers(t *testing.T) {
	Convey("Path helpers should work correctly", t, func() {
		Convey("Abs path should handle GOPATH", func() {
			path, err := absPath("github.com/coreos/etcd", "/home/user/gopath")
			So(err, ShouldBeNil)
			So(path, ShouldEqual, "/home/user/gopath/src/github.com/coreos/etcd")
		})
		Convey("NormalizePkgName should normalize names correctly", func() {
			path := "/home/user/gopath/src/github.com/coreos/etcd"
			name := normalizePkgName(".", path, "/home/user/gopath")
			So(name, ShouldEqual, "github.com/coreos/etcd")
		})
	})
}
