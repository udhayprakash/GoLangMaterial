package calc

import (
	"testing"
)

func TestAdd2(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 1, expectedOutcome: 2},
		{a: 1, b: -1, expectedOutcome: 0},
		{a: -1, b: 1, expectedOutcome: 0},
		{a: -1, b: -1, expectedOutcome: -2},
	}

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("Got: %d But expected %d", result, test.expectedOutcome)
			t.Fail()
		}
	}

}
