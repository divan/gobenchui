package main

type Benchmark string

func RunBenchmarks(vcs VCS) (chan Benchmark, error) {
	ch := make(chan Benchmark)

	go func() {
		out, err := Run(vcs.Path(), "go", "test", "-test.bench", ".")
		if err != nil {
			return
		}
		ch <- Benchmark(out)
	}()

	return ch, nil
}
