package main

import "fmt"

func main() {
	// days := 2
	// switch days {

	switch days := 1; days {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		// fallthrough
	case 4:
		fmt.Println("Four")
		fallthrough
	default:
		fmt.Println("Default case")
	}
}

// NOTE: ;(semicolon) is statement separator
// All statements after a match using
// fallthrough statement will be executed.