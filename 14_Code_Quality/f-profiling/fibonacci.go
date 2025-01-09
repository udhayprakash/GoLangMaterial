package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

// Fibonacci --- 0, 1, 1, 2, 3, 5, ....

// Fibonacci function
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	// Create a CPU profile file
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Calculate Fibonacci number
	n := 40
	result := Fibonacci(n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, result)

}


//  $ go tool pprof -png cpu.prof > profile.png
