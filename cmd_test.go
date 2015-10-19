package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestErrorGuess(t *testing.T) {
	Convey("Error guessing should guess error types correctly", t, func() {
		Convey("Panic", func() {
			typ := guessErrType(nil, panicStderr1)
			So(typ, ShouldEqual, PanicErr)
			typ = guessErrType(nil, panicStderr2)
			So(typ, ShouldEqual, PanicErr)
		})
	})
}

var panicStderr1 = `panic: nil

goroutine 33 [running]:
pkg.panicWrapper()
	/var/folders/qp/6bvmky410dn8p1yhn3b19yxr0000gn/T/gobenchui386989384/src/pkg/main.go:106 +0x23
pkg.BenchmarkStrconvConcat(0x820362100)
	/var/folders/qp/6bvmky410dn8p1yhn3b19yxr0000gn/T/gobenchui386989384/src/pkg/main_test.go:19 +0x214
testing.(*B).runN(0x820362100, 0x1)
	/usr/local/go/src/testing/benchmark.go:124 +0x9a
testing.(*B).launch(0x820362100)
	/usr/local/go/src/testing/benchmark.go:199 +0x63
created by testing.(*B).run
	/usr/local/go/src/testing/benchmark.go:179 +0x54

goroutine 1 [runnable]:
testing.(*B).run(0x820362100, 0x0, 0x0, 0x0, 0x0, 0x0)
	/usr/local/go/src/testing/benchmark.go:180 +0x7b
testing.RunBenchmarks(0x1cacc0, 0x2527a0, 0x3, 0x3)
	/usr/local/go/src/testing/benchmark.go:332 +0x75f
testing.(*M).Run(0x8202ebef8, 0x8202ba4f0)
	/usr/local/go/src/testing/testing.go:503 +0x1b8
main.main()
	pkg/_test/_testmain.go:58 +0x116`

var panicStderr2 = `testing: warning: no tests to run
panic: nil

goroutine 33 [running]:
pkg.panicWrapper()
	/var/folders/qp/6bvmky410dn8p1yhn3b19yxr0000gn/T/gobenchui386989384/src/pkg/main.go:106 +0x23
pkg.BenchmarkStrconvConcat(0x820362100)
	/var/folders/qp/6bvmky410dn8p1yhn3b19yxr0000gn/T/gobenchui386989384/src/pkg/main_test.go:19 +0x214
testing.(*B).runN(0x820362100, 0x1)`
