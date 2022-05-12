package main

import "fmt"

func main() {
	var val int = 42
	var valPtr *int = &val

	fmt.Println("val= ", val, "*valPtr=", *valPtr) // value
	fmt.Println("valPtr= ", valPtr, "&val=", &val) // address location

	// updating value directly
	val = 66
	fmt.Println("\nval= ", val, "*valPtr=", *valPtr)
	fmt.Println("valPtr= ", valPtr, "&val=", &val)

	// Indirectly updating value, via pointer
	*valPtr = 999
	fmt.Println("\nval= ", val, "*valPtr=", *valPtr)
	fmt.Println("valPtr= ", valPtr, "&val=", &val)

	// using values via pointer
	fmt.Println("val + 10    = ", val+10)     // 1009
	fmt.Println("*valPtr + 10= ", *valPtr+10) // 1009

	// updating values ia pointer
	*valPtr += 1
	fmt.Println("\nval= ", val, "*valPtr=", *valPtr) // 1000
	fmt.Println("valPtr= ", valPtr, "&val=", &val)   // 1000

	var val2 int = 55
	var valPtr2 *int = &val2

	fmt.Println("val + val2         = ", val+val2)         // 1054
	fmt.Println("*valPtr + *valPtr2 = ", *valPtr+*valPtr2) // 1054

	// fmt.Println(valPtr + valPtr2)
	// invalid operation: valPtr + valPtr2 (operator + not defined on pointer)

}

// NOTE - pointer arithmetics is not supported in go;
// But can be done using "unsafe" package

// Assignment - If two pointer are refering to same variable,
// try arithmetic on those pointers
