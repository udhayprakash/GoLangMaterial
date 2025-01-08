package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// TestDivParallel tests the Div function in parallel.
func TestDivParallel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b int
		want int
		skip bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true}, // This test should be skipped
		{7, 3, 2, false},
		{-10, 2, -5, false},
	}

	var results []string

	for _, tt := range tests {
		if tt.skip {
			results = append(results, "Skipping test with denominator 0")
			continue
		}
		result, err := Div(tt.a, tt.b)
		if err != nil {
			results = append(results, "Div("+fmt.Sprintf("%d, %d", tt.a, tt.b)+") unexpected error: "+err.Error())
		} else if result != tt.want {
			results = append(results, "Div("+fmt.Sprintf("%d, %d", tt.a, tt.b)+") = "+fmt.Sprintf("%d", result)+"; want "+fmt.Sprintf("%d", tt.want))
		} else {
			results = append(results, "Div("+fmt.Sprintf("%d, %d", tt.a, tt.b)+") = "+fmt.Sprintf("%d", result))
		}
	}

	// Read the golden file
	goldenFile := "testdata/golden.txt"
	expectedResults, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	// Compare the results
	actualResults := strings.Join(results, "\n")
	if string(expectedResults) != actualResults {
		t.Errorf("Test results do not match the golden file.\nExpected:\n%s\nActual:\n%s", expectedResults, actualResults)
	}
}
