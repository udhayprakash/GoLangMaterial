package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

// go test -v
// go test -cover
// go test -cover -coverprofile=c.out
// go tool cover -html=c.out -o coverage.html
