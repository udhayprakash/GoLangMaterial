# Testing in Go

### Benifits
1) Prevent bugs: 
    - Unit tests can help to prevent bugs by verifying the correctness of code before it is deployed.
2) Improve quality:
    - Unit tests can help to improve the quality of code by ensuring that it meets the requirements.
3) Save time:
    - Unit tests can save time by catching bugs early in the development process.
4) Increase confidence:
    - Unit tests can increase confidence in code by providing a way to verify its correctness.
    
### Rules
    1) It needs to be in a file with a name like xxx_test.go
    2) The test function must start with the word Test
    3) The test function takes one argument only t *testing.T

    To run all test scripts in current directory ,
        go test ./...

    To check the test coverage,
        Step 1: go test -coverprofile cover.out
        Step 2: go tool cover -html=cover.out
                It opens cover.out in your default browser
        Step 3:
            go test -v -coverage=./.. -coverprofile=coverage.out - json ./.. >test-report.json

## BenchMarking
    - Benchmarking is a way to measure the performance of the unit tests.
    - It is useful for identifying potential performance bottlenecks in the code.

    1) benchmarking function definition as
            func BenchmarkXxx(*testing.B)
    2) go test -bench
        go test -cpu

## Mocking Frameworks

	1) Gomock 		
    	- It is the most common and popular mocking framework to mock interfaces.
        - https://pkg.go.dev/github.com/golang/mock/gomock
	2) Go sqlmock
        - DATA-DOG sqlmock is one of the common frameworks to mock Golang’s SQL.
        - https://pkg.go.dev/github.com/data-dog/go-sqlmock
	3) Testify
        - This library is also a great alternative to mock objects and HTTP activities.
		- I use this library for its great assertion tool. 
        - https://pkg.go.dev/github.com/stretchr/testify


TODO
https://dev.to/quii/learn-go-by-writing-tests--m63


-race is the flg need to be added in command line to find data race



unit testing 
-------------
A way of testing where the behaviour of a unit of a software is tested to determine if they work properly
and exhibit the expected behaviour. 
A test is not unit test if:
- It talks to the database 
- It communicates across the network 
- It touches the file system 
- It can't run at the same time as any of your other unit tests 
- You have to do special things to your environment to run it 


Unit tests
	- Measures the quality of the code
	- Helps in verifying the small changes (refactoring, debugging) quickly
	- Helps in understanding the cause of failure quickly 
	- Helps reviewer to understand the fixes/changes

 
Table Based Testing for unit testing 
BDT - function al testing - Gingo
 
 
Table Based Testing (TBT)
 - Map/slice is used to build inputs and expected outputs 
 - Easy to implement 
 
BDT 
 - HUman readable description of software requirements
 - Requires planning and lots of effort 
 - Good in understanding the functionality of project and explaining it to stakeholders 
 - Learning Curve is little high initialy 
 - Ginkgo is BDT library 
 - Gomega is matched library 
 