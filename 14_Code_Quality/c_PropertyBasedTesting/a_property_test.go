package main

import (
	"testing"
	"testing/quick"
)

func Add(n1, n2 int) int {
	return n1 + n2
}

// 1. Commutative Property
// The order of operands in addition does not affect the result.
func TestAddCommutative(t *testing.T) {
	f := func(a, b int) bool {
		return Add(a, b) == Add(b, a) // test case
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// 2. Associative Property
// The grouping of operands in addition does not affect the result.
func TestAddAssociative(t *testing.T) {
	f := func(a, b, c int) bool {
		return Add(Add(a, b), c) == Add(a, Add(b, c)) // test case
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// 3. Identity Property
// Multiplying a number by 1 leaves it unchanged.

func Multiply(n1, n2 int) int {
	return n1 * n2
}

func TestMultiplyIdentity(t *testing.T) {
	f := func(a int) bool {
		return Multiply(a, 1) == a // test case
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// 4. Reversing a String Twice
// should return the original string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestReverseTwice(t *testing.T) {
	f := func(s string) bool {
		return Reverse(Reverse(s)) == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// 5. Concatenation of Slices
// Total final length should be equal to original
func TestSliceConcatenation(t *testing.T) {
	f := func(a, b []int) bool {
		return len(append(a, b...)) == len(a)+len(b)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
