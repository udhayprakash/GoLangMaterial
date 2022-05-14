package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)

	// blocking infinitely - NOT recommended
	var input string
	fmt.Scanln(&input)

	fmt.Println("Done")
}
