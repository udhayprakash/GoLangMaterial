package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // shadowing with single assignment
		fmt.Println(x)
	}
	fmt.Println(x)

	if x > 5 {
		// shadowing with multiple assignments
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	// shadowing package names
	// fmt := "oops"
	// fmt.Println(fmt)
}

/*
- some times, these shadow variables can introduce subtle bugs
- Neither 'go vet' or 'golangci-lint' can detect these.

- shadow tool
	- go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	- USAGE
			shadow ./..
*/
