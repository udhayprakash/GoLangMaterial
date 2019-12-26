package main

import (
	"fmt"
)
/*Purpose: panic
	- To raise an error to fail fast when something went
      unexpectedly wrong.

	- Note that unlike some languages which use exceptions
      for handling of many errors, in Go it is idiomatic to
      use error-indicating return values wherever possible.


*/
func main() {
	fmt.Println("first Statement")
	//panic("a problem")
	fmt.Println("Last Statement")

}