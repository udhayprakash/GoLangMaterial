package main

import "testing"

func Add(n1, n2 int) int {
	return n1 + n2
}

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		result := Add(a, b)
		if result != a+b {
			t.Errorf("Add(%d, %d) = %d; expected %d", a, b, result, a+b)
		}
	})
}

// Fuzz tests are designed to run indefinitely unless limited by -fuzztime
// go test -fuzz=FuzzAdd

// Limit by Time Duration
// go test -fuzz=FuzzAdd -fuzztime=5s


// limit By number of Executions (100 iterations times)
// go test -fuzz=FuzzAdd -fuzztime=100x


// Combine Flags
// go test -fuzz=FuzzAdd -fuzztime=10s -v


// Stop on First Failure
// go test -fuzz=FuzzAdd -fuzztime=1s
