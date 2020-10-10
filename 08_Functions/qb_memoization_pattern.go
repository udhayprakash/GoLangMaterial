package main

import "fmt"

/*
Purpose: Implementing Memoization Design pattern
*/
// fibonacci series - 0, 1, 1, 2, 3, 5, 8, .....

var fibValues = make(map[int]int)

func fibonacciNum(num int) int {
	val, isKeyPresent := fibValues[num]
	if isKeyPresent != false{
		fmt.Printf("\tfibValues[%2d]=%4d\n", num, val)
		return val
	}
	fmt.Printf("in func - fibonacciNum(%d)\n", num)
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacciNum(num-1) + fibonacciNum(num-2)
	}
}

func main() {
	//fmt.Println("fibonacciNum(4)=", fibonacciNum(4))

	for i := 0; i < 15; i++ {
		fibValues[i]  = fibonacciNum(i)
		fmt.Printf("fibonacci(%d) :%2d\n", i, fibValues[i])

	}
	fmt.Println("fib_values:", fibValues)

}

/*
in func - fibonacciNum(4)
in func - fibonacciNum(3)
in func - fibonacciNum(2)
in func - fibonacciNum(1)
in func - fibonacciNum(0)
in func - fibonacciNum(1)
in func - fibonacciNum(2)
in func - fibonacciNum(1)
in func - fibonacciNum(0)
fibonacciNum(4)= 3

*/