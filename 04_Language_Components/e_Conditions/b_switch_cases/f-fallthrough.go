package main

import "fmt"

func main() {
	// days := 4
	// switch days {

	switch days := 4; days {
	case 1:
		fmt.Println("One")
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
/*
NOTE:
	1) ;(semicolon) is statement separator
	2) All statements after a match using
     fallthrough statement will be executed.
*/ 