package main

import (
	"bytes"	"fmt"
)


func main() {
	var n1 = []byte{1, 2, 3}
	var n2 = []byte{1, 2, 3}
	var n3 = []byte{3, 2, 1}

	fmt.Println("bytes.Compare(n1, n2) 		= ", bytes.Compare(n1, n2)) // 0
	fmt.Println("reflect.DeepEqual(n1, n2) 	= ", reflect.DeepEqual(n1, n2)) // true
	
	fmt.Println("bytes.Compare(n1, n3) 		= ", bytes.Compare(n1, n3)) // -1
	
}
