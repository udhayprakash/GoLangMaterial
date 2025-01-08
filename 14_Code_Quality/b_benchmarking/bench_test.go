package benchmarking

import (
	"testing"
	"time"
)

func BenchmarkEfficient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Efficient(1, 2)
	}
}

func BenchmarkParseRFC3339(b *testing.B) {
	now := time.Now().UTC().Format(time.RFC3339Nano)

	for i := 0; i < b.N; i++ {
		if _, err := time.Parse(time.RFC3339, now); err != nil {
			b.Fatal(err)
		}
	}
}

// Usage: go test -bench=.

// > go test -run ^$ -bench BenchmarkParseRFC3339 -cpuprofile cpu.prof
// > go tool pprof cpu.prof
