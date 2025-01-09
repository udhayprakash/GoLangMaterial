package main

import (
	"testing"
	"unicode/utf8"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		reversed := Reverse(s)
		doubleReversed := Reverse(reversed)
		if s != doubleReversed {
			t.Errorf("Reverse(Reverse(%q)) = %q; expected %q", s, doubleReversed, s)
		}
		if utf8.ValidString(s) && !utf8.ValidString(reversed) {
			t.Errorf("Reverse produced invalid UTF-8 string: %q", reversed)
		}
	})
}
