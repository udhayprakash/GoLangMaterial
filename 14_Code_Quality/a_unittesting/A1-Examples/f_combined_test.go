package Bexample

import (
	"fmt"
	"testing"
)

// go run will not work for this program, as main() function is not defined
// go test -v
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMin(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}

	ans = IntMin(2, -2)
	if ans != 2 {
		// t.Fatalf("IntMin(2, -2) = %d; want -2", ans)
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
	// t.Fatalf() to stop the test immediately if a critical failure occurs
}

func TestIntMinTableDriven(t *testing.T) {
	t.Parallel() //  To run tests in parallel
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 1, 1},
		{1, 0, 0},
		{-2, -3, -2},
		{-2, -3, -3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// Subtests with t.Run()
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Multiple subtests
func TestSubtests(t *testing.T) {
	t.Run("Subtest 1", func(t *testing.T) {
		result := IntMin(2, 3)
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})

	t.Run("Subtest 2", func(t *testing.T) {
		result := IntMin(-1, 0)
		if result != -1 {
			t.Errorf("Expected -1, got %d", result)
		}
	})
}

/*
$ go test -v f_combined_test.go
=== RUN   TestIntMin
    f_combined_test.go:25: IntMin(2, -2) = -2; want -2
--- FAIL: TestIntMin (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/-2,-3
    f_combined_test.go:46: got -3, want -2
=== RUN   TestIntMinTableDriven/-2,-3#01
--- FAIL: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- FAIL: TestIntMinTableDriven/-2,-3 (0.00s)
    --- PASS: TestIntMinTableDriven/-2,-3#01 (0.00s)
FAIL
FAIL    command-line-arguments  0.002s
FAIL

*/