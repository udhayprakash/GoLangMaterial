package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(true, false)

	val_1 := true
	fmt.Println("reflect.TypeOf(val_1) =", reflect.TypeOf(val_1))

	var val_2 bool        // default bool value is false
	fmt.Println("reflect.TypeOf(val_2) =", reflect.TypeOf(val_2))
	fmt.Println("val_2 =", val_2)  

}
