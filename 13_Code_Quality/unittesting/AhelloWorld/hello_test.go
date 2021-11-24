package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	expecting := "Hello World"

	if got != expecting {
		t.Errorf("Got %s \t Expecting:%s", got, expecting)
	}
}
