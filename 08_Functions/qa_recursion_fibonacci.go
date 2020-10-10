package main

import "fmt"

// fibonacci series - 0, 1, 1, 2, 3, 5, 8, .....

func fibonacciNumber(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacciNumber(num - 1) + fibonacciNumber(num -2)
	}
}

func main()  {
	fmt.Println("fibonacciNumber(3):", fibonacciNumber(3))
	fmt.Println("fibonacciNumber(5):", fibonacciNumber(5))

	for i := 0; i < 25; i++{
		fmt.Printf("fibonacciNumber(%2d) = %4d\n", i, fibonacciNumber(i))
	}
}