package bench

import "testing"

func BenchmarkEfficient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Efficient(1, 2)
	}
}

// Usage: go test -bench=.
