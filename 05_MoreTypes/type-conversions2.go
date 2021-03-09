package main

import "fmt"

func main() {
	s := "Go"
	bs := []byte(s)
	fmt.Printf("%s\n", bs) // Output: Go
	fmt.Printf("%d\n", bs) // Output: [71 111]

	bs2 := []byte{71, 111}
	fmt.Printf("%s\n", bs2) // Output: Go
}

/*
NOTE:
	[ ]byte: this is a byte slice i.e. a dynamic size array
	       that contains bytes i.e. each element is a UTF-8
		   character.
	String: read-only slices of bytes i.e. immutable
*/
