Testing in Go
=============

    Rules
    -----
    1) It needs to be in a file with a name like xxx_test.go
    2) The test function must start with the word Test
    3) The test function takes one argument only t *testing.T

    To run all test scripts in current directory , 
        go test ./...
    
    To check the test coverage,
        Step 1: go test -coverprofile cover.out
        Step 2: go tool cover -html=cover.out 
                It opens cover.out in your default browser


BenchMarking
-------------

    1) benchmarking function definition as 
            func BenchmarkXxx(*testing.B)
    2) go test -bench
    go test -cpu

TODO
https://dev.to/quii/learn-go-by-writing-tests--m63  