package main

import (
	"testing"
	"testing/quick"
)

// Square returns the square of an integer.
func Square(x int) int {
	return x * x
}

// IsEven checks if a number is even.
func IsEven(x int) bool {
	return x%2 == 0
}

// Concat concatenates two strings.
func Concat(a, b string) string {
	return a + b
}

// DoubleSlice doubles each element in a slice of integers.
func DoubleSlice(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = v * 2
	}
	return result
}


// TestSquareNonNegative tests that the square of any number is non-negative.
func TestSquareNonNegative(t *testing.T) {
	f := func(x int) bool {
		return Square(x) >= 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestEvenSquare tests that the square of an even number is also even.
func TestEvenSquare(t *testing.T) {
	f := func(x int) bool {
		if IsEven(x) {
			return IsEven(Square(x))
		}
		return true // Skip odd numbers
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestConcatLength tests that the length of concatenated strings is the sum of their lengths.
func TestConcatLength(t *testing.T) {
	f := func(a, b string) bool {
		return len(Concat(a, b)) == len(a)+len(b)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestDoubleSliceLength tests that doubling a slice does not change its length.
func TestDoubleSliceLength(t *testing.T) {
	f := func(s []int) bool {
		return len(DoubleSlice(s)) == len(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestDoubleSliceValues tests that each element in the doubled slice is twice the original.
func TestDoubleSliceValues(t *testing.T) {
	f := func(s []int) bool {
		doubled := DoubleSlice(s)
		for i, v := range s {
			if doubled[i] != v*2 {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestReverseConcat tests that reversing the concatenation of two strings is the same as
// concatenating their reverses in reverse order.
func TestReverseConcat(t *testing.T) {
	f := func(a, b string) bool {
		concat := Concat(a, b)
		reversedConcat := Reverse(concat)
		reversedA := Reverse(a)
		reversedB := Reverse(b)
		expected := Concat(reversedB, reversedA)
		return reversedConcat == expected
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
