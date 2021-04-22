package main

import "fmt"

func main() {
	var val int = 42
	var valPtr *int = &val

	fmt.Println("val= ", val, "*valPtr=", *valPtr)
	fmt.Println("valPtr= ", valPtr, "&val=", &val)

	// updating value directly
	val = 66
	fmt.Println("\nval= ", val, "*valPtr=", *valPtr)
	fmt.Println("valPtr= ", valPtr, "&val=", &val)

	// Indirectly updating value, via pointer
	*valPtr = 999
	fmt.Println("\nval= ", val, "*valPtr=", *valPtr)
	fmt.Println("valPtr= ", valPtr, "&val=", &val)

	fmt.Println("val + 10= ", val+10)
	fmt.Println("*valPtr + 10= ", *valPtr+10)

	var val2 int = 55
	var valPtr2 *int = &val2

	fmt.Println("val + val2         = ", val+val2)
	fmt.Println("*valPtr + *valPtr2 = ", *valPtr+*valPtr2)

}

// NOTE - pointer arithmetics is not supported in go;
// But can be done using "unsafe" package

// Assignment - if two pointer are refering to same variable,
// try arithmetic on those pointers
