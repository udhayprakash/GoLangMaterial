package main

import "fmt"

func main() {
	var X, Y int

	fmt.Printf("\nInitially, X: %3d, Y: %d", X, Y)

	// Sscan scans the argument string, storing successive space-separated values into successive arguments. Newlines count as space.
	fmt.Sscan("100\n200", &X, &Y)
	fmt.Printf("\nSscan,     X: %3d, Y: %d", X, Y)

	// Sscanln is similar to Sscan, but stops scanning at a newline and after the final item there must be a newline or EOF.
	fmt.Sscanln("300\n400", &X, &Y)
	fmt.Printf("\nSscanln,   X: %3d, Y: %d", X, Y)
	// Observe Y value

	// Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format.
	fmt.Sscanf("(500, 600)", "(%d, %d)", &X, &Y)
	fmt.Printf("\nSscanf,    X: %3d, Y: %d", X, Y)

}
