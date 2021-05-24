package main

import "fmt"

/*Purpose: panic
	- To raise an error to fail fast when something went
      unexpectedly wrong.

	- Note that unlike some languages which use exceptions
      for handling of many errors, in Go it is idiomatic to
      use error-indicating return values wherever possible.

	- Go’s type system catches many mistakes at compile time,
	but others, like an out-of-bounds array access or nil pointer
    dereference, require checks at run time. When the Go runtime
	detects these mistakes, it pani cs.
*/
func main() {
	a, b := 1, 0
	ans := a / b

	fmt.Println(ans)

}

// panic: runtime error: integer divide by zero
