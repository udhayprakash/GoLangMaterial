package main

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Log("Testing First Case")
	t.Logf("Testing First Case")
	fmt.Println("This is a print statement")
}

/*
OUTPUT:
-------

$ go run a_simple_test.go 
go: cannot run *_test.go files (a_simple_test.go)

$ go test
PASS
ok      /GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples    0.001s

$ go test -v
=== RUN   TestFirst
    a_simple_test.go:6: Testing First Case
    a_simple_test.go:7: Testing First Case
--- PASS: TestFirst (0.00s)
PASS
ok      /GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples    0.001s

$ go test -v
=== RUN   TestFirst
    a_simple_test.go:9: Testing First Case
    a_simple_test.go:10: Testing First Case
This is a print statement
--- PASS: TestFirst (0.00s)
PASS
ok      /GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples    0.002s

$ go test -test.v
=== RUN   TestFirst
    a_simple_test.go:9: Testing First Case
    a_simple_test.go:10: Testing First Case
This is a print statement
--- PASS: TestFirst (0.00s)
PASS
ok      /GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples    0.001s


*/