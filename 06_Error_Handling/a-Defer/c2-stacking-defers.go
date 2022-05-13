package main

import (
	"fmt"
)

func main() {
	fmt.Println("main()   - starting")

	defer fmt.Println("defer main()")
	myFunc()
	AnotherFunc()

	fmt.Println("main()   - ending")
}

func myFunc() {
	fmt.Println("\tmyFunc() - starting")

	defer fmt.Println("\t\tdefer statement1")

	fmt.Println("\tmyFunc() - ending")
}

func AnotherFunc() {
	fmt.Println("\tAnotherFunc() - starting")

	defer fmt.Println("\t\tdefer statement2")

	fmt.Println("\tAnotherFunc() - ending")
}


/*
OUTPUT:
-------
	main()   - starting
			myFunc() - starting
			myFunc() - ending
					defer statement1
			AnotherFunc() - starting
			AnotherFunc() - ending
					defer statement2
	main()   - ending
	defer main()

*/
