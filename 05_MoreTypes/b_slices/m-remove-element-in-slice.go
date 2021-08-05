package main

import "fmt"

func main() {
	// Method 1 - fast, but changes order : O(1) - constant time
	a := []string{"A", "B", "C", "D", "E"}
	//             0    1    2    3    4
	i := 2

	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.
	fmt.Println(a)     // [A B E D]

	// Method 2 - slow, but retains order : O(n) - linear time
	a = []string{"A", "B", "C", "D", "E"}
	//             0    1    2    3    4
	i = 2
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	fmt.Println(a)       // [A B D E]

	// Method 3 - Simple and clean
	a = []string{"A", "B", "C", "D", "E"}
	i = 2
	a = append(a[:i], a[i+1:len(a)]...)
	fmt.Println(a) // [A B D E]
}
