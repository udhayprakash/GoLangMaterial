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
	input := []int{1, 2, 3, 4, 5}
	actual := Max(input)
	expected := 6

	if actual != expected { // FAILURE, but will pass
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

func TestMax2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	actual := Max(input)
	expected := 6

	if actual != expected { // FAILURE, will FAIL
		t.Fail()
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

/*
OUTPUT:
-go test -v c_example_test.go
=== RUN   TestMax
    c_example_test.go:23: Expected 6, got 5
--- PASS: TestMax (0.00s)
=== RUN   TestMax2
    c_example_test.go:34: Expected 6, got 5
--- FAIL: TestMax2 (0.00s)
FAIL
FAIL    command-line-arguments  0.023s
FAIL
*/
