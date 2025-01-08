package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}


// PreRequisite for this script, is to have testdata/goldenfile.txt will test cases defined
func TestDivGoldenFileBased(t *testing.T) {
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
	var results []string
	for _, tt := range testCases {
		if tt.skip {
			results = append(results, "Skipping test with denominator 0")
			continue
		}
		result, err := Div(tt.a, tt.b)
		if err != nil {
			// t.Errorf("Div(%d, %d) unexpected error: %v", tt.a, tt.b, err)
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
