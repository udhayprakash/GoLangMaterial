package main

import "testing"

func TestFoo(t *testing.T) {
	t.Logf("Testing Foo")
}

/*
OUTPUT:
-------
	~go test
	PASS
	ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/A-Examples      0.034s

	~go test -v
	=== RUN   TestFoo    a_example_test.go:6: Testing Foo
	--- PASS: TestFoo (0.00s)
	PASS
	ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/A-Examples      0.041s

	~go test -test.v
	=== RUN   TestFoo
		a_example_test.go:6: Testing Foo
	--- PASS: TestFoo (0.00s)
	PASS
	ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/A-Examples      0.032s



*/
