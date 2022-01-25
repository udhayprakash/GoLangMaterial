package main

import "fmt"

func counter(existingValue int) (newValue int) {
	newValue = existingValue + 1
	return newValue
}

func counter2(valPtr *int) {
	*valPtr++
}
func main() {
	// Method 1
	var number int = 0
	number = counter(number)
	fmt.Println("After counter(number), number =", number) // 1

	number = counter(number)
	fmt.Println("After counter(number), number =", number) // 2
	fmt.Println()

	// Method 2
	var number2 int = 0
	counter2(&number2)
	fmt.Println("After counter2(&number2), number2 =", number2) // 1

	counter2(&number2)
	fmt.Println("After counter2(&number2), number2 =", number2) // 2
	fmt.Println()

	// Method 3
	counter := NewCounter()

	counter() // n:1
	counter() // n:2
	counter() // n:3

	// fmt.Println(n) // undefined: n

	otherCounter := NewCounter()
	otherCounter() // n : 1
	otherCounter() // n : 2

	counter()      // n:4
	otherCounter() // n : 3
}

// function Literal
// Count prints the number of times it has been invoked.
func NewCounter() (count func()) {
	n := 0
	return func() {
		n++
		fmt.Println("n:", n)
	}
}
