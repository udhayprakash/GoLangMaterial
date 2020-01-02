package main

import (
"fmt"
)

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

	// prints "1"
	fmt.Println(someCounter())

	// prints "4"
	// because someCounter() increases i to 2
	// then doubleUp() doubles what was returned
	fmt.Println(doubleUp(someCounter()))

	// prints "3"
	// because the previous doubleUp() didn't affect i
	// only the value returned from someCounter
	fmt.Println(someCounter())
}