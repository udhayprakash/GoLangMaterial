package main

import "fmt"

// return results of A + B and A * B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func SumAndProduct1(A, B int) (add int, multiplied int) {
	add = A + B
	multiplied = A * B
	return
}

// Since return arguments are named, the function automatically
// returns them
func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)

	xPLUSy, xTIMESy = SumAndProduct1(x, y)

	fmt.Printf("\n\n%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
