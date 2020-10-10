package main

import "fmt"

/*
Purpose: Implementing Memoization Design pattern
*/
var fibValues = make(map[int]int)

func fibonacciNum(num int) int {
	val, isKeyPresent := fibValues[num]
	if isKeyPresent != false {
		fmt.Println("\tfibValues[num]", num, val)
		return val
	} else if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacciNum(num-1) + fibonacciNum(num-2)
	}
}
func main() {

	for i := 0; i < 10; i++ {
		fibValues[i] = fibonacciNum(i)
		fmt.Printf("fibonacci(%d) :%2d\n", i, fibValues[i])
	}
	fmt.Println("fib_values", fibValues)
}

// fib_values map[0:0 1:1 2:1 3:2 4:3 5:5 6:8 7:13 8:21 9:34]
