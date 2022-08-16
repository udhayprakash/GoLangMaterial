package palindrome

import "testing"

func TestPalindromeCheck(t *testing.T) {
testTable:
	[]struct {
		inputStr       string
		expectedOutput bool
	}{
		{inputStr: "hello", expectedOutput: false},
		{inputStr: "ollo", expectedOutput: true},
		{inputStr: "ollo ", expectedOutput: false},
	}
	for _, test := range testTable {
		result := palindrome(test.inputStr)
		if result != test.expectedOutput {
			t.Logf("Got: %d But expected %d", result, test.expectedOutcome)
			t.Fail()
		}
	}
}
