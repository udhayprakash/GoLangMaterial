// Expression switch statement example
package main

import "fmt"

// All statements after a match using
// fallthrough statement will be executed.

func main() {

	// switch statement with both separated by ;
	// statement 	days:= 2
	// expression	days
	switch days := 2; days {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
		fallthrough
	default:
		fmt.Println("Default case")
	}
}
