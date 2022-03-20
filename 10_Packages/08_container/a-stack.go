package main

import (
	"container/list"
	"fmt"
	"reflect"
)

/*
STACK
	----------------------------|
->								|
<-								|
	----------------------------|

*/
func main() {
	stack := list.New()
	fmt.Printf("Type - %[1]T \n", stack)

	stack.PushFront("first")
	fmt.Println("stack = ", stack)

	stack.PushFront("second")
	fmt.Println("stack = ", stack)

	stack.PushFront("third")
	fmt.Println("stack = ", stack)

	// Get the reflect.Value fv for the unexported  field len
	fv := reflect.ValueOf(stack).Elem().FieldByName("len")
	fmt.Println(fv.Int())

}
