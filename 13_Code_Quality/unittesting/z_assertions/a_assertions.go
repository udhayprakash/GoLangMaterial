package main

import (
	"testing"
)


func TestDemo(t *testing.T){
	one := 111
	two := 111
	assert.Equal(t, one, two, "the two variables should have same value")
}