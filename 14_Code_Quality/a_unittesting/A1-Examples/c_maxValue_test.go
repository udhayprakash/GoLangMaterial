package main

import "testing"

func Max(numbers []int) int {
	var max int
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func TestMax(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	actual := Max(inputSlice)
	expected := 5

	if actual != expected {
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

func TestMax2(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	actual := Max(inputSlice)
	expected := 6

	if actual != expected { // Failure, will fail
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

/*
OUTPUT:
------
$ go test -v c_maxValue_test.go
testing: warning: no tests to run
PASS
ok      command-line-arguments  0.003s [no tests to run]
$ go test -v c_maxValue_test.go
=== RUN   TestMax
--- PASS: TestMax (0.00s)
PASS
ok      command-line-arguments  0.001s
$ go test -v c_maxValue_test.go
=== RUN   TestMax
    c_maxValue_test.go:21: Expected 50, got 5
--- PASS: TestMax (0.00s)
PASS
ok      command-line-arguments  0.001s
$ go test -v c_maxValue_test.go
=== RUN   TestMax
--- PASS: TestMax (0.00s)
PASS
ok      command-line-arguments  0.002s
$ go test -v c_maxValue_test.go
=== RUN   TestMax
--- PASS: TestMax (0.00s)
=== RUN   TestMax2
    c_maxValue_test.go:31: Expected 6, got 5
--- PASS: TestMax2 (0.00s)
PASS
ok      command-line-arguments  0.002s

*/