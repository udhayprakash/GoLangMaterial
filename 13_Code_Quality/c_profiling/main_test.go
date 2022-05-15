package main

import "testing"

func TestGetVal(t *testing.T) {
	for i := 0; i < 1000; i++ { // running it a 1000 times
		if Fib2(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
