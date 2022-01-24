package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var n1 = []byte{1, 2, 3}

	fmt.Println(len(n1), n1 == nil) // 3 false

	var n2 = []byte{1, 2, 3}
	var n3 = []byte{3, 2, 1}

	// n1 == n2 // slice can only be compared to nil

	fmt.Println("bytes.Compare(n1, n2) 		= ", bytes.Compare(n1, n2))        // 0
	fmt.Println("reflect.DeepEqual(n1, n2) 	= ", reflect.DeepEqual(n1, n2)) // true

	fmt.Println("bytes.Compare(n1, n3) 		= ", bytes.Compare(n1, n3))        // -1
	fmt.Println("reflect.DeepEqual(n1, n3) 	= ", reflect.DeepEqual(n1, n3)) // false

}

// assignment: try for slices of other data types and 2-D slices
