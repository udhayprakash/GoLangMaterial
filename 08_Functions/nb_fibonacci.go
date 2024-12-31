package main

import "fmt"

// fibonacci Series: 0 1 1 2 3 5 8 13 .....
func fibonacci() func() int {
	a, b := 0, 1        // context with be remembered
	return func() int { // anonymous function
		// a = b
		// b = a + b
		a, b = b, a+b // unpacking
		return a

	}
}

func main() {
	series1 := fibonacci()

	fmt.Println("series1() = ", series1())
	fmt.Println("series1() = ", series1())
	fmt.Println("series1() = ", series1())
	fmt.Println("series1() = ", series1())
	fmt.Println("series1() = ", series1())
	fmt.Println()

	// context separates for every primary call
	series2 := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("series2() = ", series2())

	}
}
