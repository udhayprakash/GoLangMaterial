package main

import "fmt"

func main() {
	// comparision on arrays
	a1 := [3]int{ 1, 2, 3}
	a2 := [3]int{ 2, 3, 1}
	a3 := [3]int{ 3, 1, 2}
	a4 := [3]int{ 1, 2, 3}

	fmt.Println("a1 == a2 :", a1 == a2)
	fmt.Println("a1 == a3 :", a1 == a3)
	fmt.Println("a1 == a4 :", a1 == a4)
}
