package main

import "fmt"

func fibonacciNum(num int) int {
	if num == 0  {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacciNum(num -1) + fibonacciNum(num -2)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("fibonacci(%d) :%2d\n", i, fibonacciNum(i))
	}
}
