package main

import "testing"

func TestFirst(t *testing.T) {
	t.Logf("Testing First Case")
}

/*
OUTPUT:
-------
-go test
PASS
ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/a_unittesting/A1-Examples     0.069s

-go test -v
=== RUN   TestFirst
    a_example_test.go:6: Testing First Case
--- PASS: TestFirst (0.00s)
PASS
ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/a_unittesting/A1-Examples     0.085s

-go test -test.v
=== RUN   TestFirst
    a_example_test.go:6: Testing First Case
--- PASS: TestFirst (0.00s)
PASS
ok      github.com/udhayprakash/GoLangMaterial/13_Code_Quality/a_unittesting/A1-Examples     0.039s
*/
