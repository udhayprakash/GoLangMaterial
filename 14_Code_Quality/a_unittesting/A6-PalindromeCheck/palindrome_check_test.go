package main

import "testing"

func TestPalindromeCheck(t *testing.T) {
	testTable := []struct {
		inputStr       string
		expectedOutput bool
	}{
		{inputStr: "hello", expectedOutput: false},
		{inputStr: "ollo", expectedOutput: true},
		// {inputStr: "ollo ", expectedOutput: false},
	}
	for _, test := range testTable {
		result := isPalindrome(test.inputStr)
		if result != test.expectedOutput {
			t.Logf("Got: %v But expected %v", result, test.expectedOutput)
			t.Fail()
		}
	}
}