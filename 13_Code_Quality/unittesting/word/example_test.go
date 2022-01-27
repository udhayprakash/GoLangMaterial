package word

import (
	"testing"
)

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestPalindrome(t *testing.T) {
	if IsPalindrome("detartrated") == false {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}

}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

/*
Usage:
	file name should end with _test.go

	To test all the test cases
		go test

	To test all test cases, and display detailed result,
		go test -v

	To run all the test cases, whose names match the regular expression
		go test -v -run="French|Canal"

	To run specific test case(s),
		go test -run ^TestCanalPalindrome$

	To add timeout restriction for the tests
		go test -timeout 30s
*/
