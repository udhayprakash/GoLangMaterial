package main

import "fmt"

// fallthrough will help to execute the consecutive case/default
func main() {
	// days := 2
	// switch days {

	switch days := 2; days {
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
	when 2 is succces , we  want to rn both 2 & 3
	when 3 is succesful , we want to run on 3

*/
