package main

import (
	"errors"
	"testing"
)

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestDiv(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b int
		want int
		skip bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true}, // This test should be skipped
		{7, 3, 2, false},
		{-10, 2, -5, false},
	}
	for _, tt := range testCases {
		if tt.skip {
			t.Logf("Skipping test with denominator %d", tt.b)
			continue
		}
		result, err := Div(tt.a, tt.b)
		if err != nil {
			t.Errorf("Div(%d, %d) unexpected error: %v", tt.a, tt.b, err)
		} else if result != tt.want {
			t.Errorf("Div(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.want)
		}
	}
}

/*
OUTPUT:
-------
$ go test g_skip_test.go 
--- FAIL: TestDiv (0.00s)
    g_skip_test.go:30: Skipping test with denominator 0
    g_skip_test.go:34: Div(10, 0) unexpected error: division by zero
FAIL
FAIL    command-line-arguments  0.001s
FAIL

$ go test g_skip_test.go 
ok      command-line-arguments  0.002s

$ go test g_skip_test.go  -v
=== RUN   TestDiv
=== PAUSE TestDiv
=== CONT  TestDiv
    g_skip_test.go:30: Skipping test with denominator 0
--- PASS: TestDiv (0.00s)
PASS
ok      command-line-arguments  0.002s
*/
