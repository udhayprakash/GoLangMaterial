package main

import (
	"testing"
)
func TestHello(t *testing.T) {
	got := Hello()
	expecting := "Hello world"

	if got != expecting {
		t.Errorf("Got %s \t Expecting:%s", got, expecting)
	}
}
