package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int32 = 42
	var b string = "forty two"

	valueOfA := reflect.ValueOf(a)
	fmt.Println("a =", a, "\t\tvalueOfA=", valueOfA)
	fmt.Println("valueOfA.Interface()=", valueOfA.Interface())

	valueOfB := reflect.ValueOf(b)
	fmt.Println("\n\nb=", b, "\tvalueOfB=", valueOfB)
	fmt.Println("valueOfB.Interface()=", valueOfB.Interface())
}
