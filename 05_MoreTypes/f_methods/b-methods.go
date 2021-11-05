package main

import (
	"fmt"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

// Ordinary function
func getFullName(first, mid, last string) string {
	return first + " " + mid + " " + last
}

// Method
func (n Name) FullName() string {
	return n.First + " " + n.Middle + " " + n.Last
}

// Method
func (n Name) Length() int {
	// return len(n.First + " " + n.Middle + " " + n.Last)
	return len(n.FullName()) // -- accessing another method, from this method
}

func main() {
	n1 := Name{"Udhay", "Prakash", "Pethakamsetty"}
	fmt.Println("n1 = ", n1)

	// Method 1 - using functions
	fullname := getFullName(n1.First, n1.Middle, n1.Last)
	fmt.Println("fullname      =", fullname)

	//  Method 2 - using method
	fmt.Println("n1.FullName() =", n1.FullName())
	fmt.Printf("n1.Length()   =%v\n", n1.Length())

}
 