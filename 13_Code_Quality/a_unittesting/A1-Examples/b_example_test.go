package main

import "testing"

func TestSecond(t *testing.T) {
	t.Logf("Testing Second")
	t.Fail()
}

func TestThird(t *testing.T) {
	t.Logf("Testing Foo")
	t.Fail()
	t.Logf("Another log from Foo")
}

/*
OUTPUT:
-go test
--- FAIL: TestSecond (0.00s)
    b_example_test.go:6: Testing Second
--- FAIL: TestThird (0.00s)
    b_example_test.go:11: Testing Foo
    b_example_test.go:13: Another log from Foo
FAIL
exit status 1
FAIL    github.com/udhayprakash/GoLangMaterial/13_Code_Quality/a_unittesting/A1-Examples     0.102s
*/
