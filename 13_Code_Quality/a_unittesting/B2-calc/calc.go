package calc

import "log"

var DOZEN int = 12

// Add a and b
func Add(a, b int) int {
	return a + b
}

// Subtract b from a.
func Subtract(a, b int) int {
	return a - b
}

// Multiply a and b.
func Multiply(a, b int) int {
	return a * b
}

// Divide a by b.
func Divide(a, b int) int {
	return a / b
}

func dozenMeans() {
	log.Println("DOZEN:", DOZEN)
}
