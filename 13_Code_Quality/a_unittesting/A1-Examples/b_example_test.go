package main

import "testing"

func TestMax(t *testing.T) {
	t.Logf("Testing Foo")
	t.Fail()
}

func TestMax2(t *testing.T) {
	t.Logf("Testing Foo")
	t.Fail()
	t.Logf("Another log from Foo")
}

/*
OUTPUT:

	~go test
	--- FAIL: TestMax (0.00s)
		b_example_test.go:6: Testing Foo--- FAIL: TestMax2 (0.00s)
		b_example_test.go:11: Testing Foo
		b_example_test.go:13: Another log from FooFAILexit status 1
	FAIL    github.com/udhayprakash/GoLangMaterial/13_Code_Quality/A-Examples      0.025s

*/
