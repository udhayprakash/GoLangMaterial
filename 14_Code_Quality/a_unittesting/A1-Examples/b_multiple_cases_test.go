package main

import "testing"

func TestSecond(t *testing.T) {
	t.Logf("Testing Second")
	t.Fail()
}

func TestThird(t *testing.T) {
	t.Log("Testing Third")
	t.Fail()
	t.Logf("Another log from Third")
}


/*
OUTPUT:
-------
 $ go test -v
=== RUN   TestFirst
    a_simple_test.go:9: Testing First Case
    a_simple_test.go:10: Testing First Case
This is a print statement
--- PASS: TestFirst (0.00s)
=== RUN   TestSecond
    b_multiple_cases_test.go:6: Testing Second
--- FAIL: TestSecond (0.00s)
=== RUN   TestThird
    b_multiple_cases_test.go:11: Testing Third
    b_multiple_cases_test.go:13: Another log from Third
--- FAIL: TestThird (0.00s)
FAIL
exit status 1
FAIL    /GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples    0.001s


 $ go test -v b_multiple_cases_test.go 
=== RUN   TestSecond
    b_multiple_cases_test.go:6: Testing Second
--- FAIL: TestSecond (0.00s)
=== RUN   TestThird
    b_multiple_cases_test.go:11: Testing Third
    b_multiple_cases_test.go:13: Another log from Third
--- FAIL: TestThird (0.00s)
FAIL
FAIL    command-line-arguments  0.001s
FAIL

*/