package benchmarking

import "time"

// Efficient is a super efficient function which needs a benchmark.
func Efficient(a, b int) int {
	time.Sleep(500 * time.Nanosecond)
	return a + b
}
