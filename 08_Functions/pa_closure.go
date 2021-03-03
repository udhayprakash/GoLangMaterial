package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func doubleUp(current int) int {
	currentCount := current
	return currentCount * 2
}
func main() {
	someCounter := counter()
	fmt.Println("someCounter  :", someCounter)
	fmt.Println("someCounter():", someCounter()) // 1

	fmt.Println("doubleUp:", doubleUp)
	fmt.Println("doubleUp(someCounter()):", doubleUp(someCounter())) // 4

	fmt.Println("someCounter():", someCounter()) // 3
	// because the previous doubleUp() didn't affect i
	// only the value returned from someCounter

}
