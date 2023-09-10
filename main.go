package gofbench

import (
	"time"
)

func Benchmark(fn func()) (start time.Time, duration time.Duration) {
	defer func() {
		duration = time.Since(start)
	}()
	start = time.Now()
	fn()
	return start, duration
}
