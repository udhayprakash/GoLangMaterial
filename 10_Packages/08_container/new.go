package main 

import (
	"container/list"
	"fmt"
	"reflect"
)

func  main()  {
	myList := list.New()
	myList.PushFront("foo")
	myList.PushFront("bar")

	// Get the reflect.Value fv for the unexported  field len
	fv  := reflect.ValueOf(myList).Elem().FieldByName("len")
	fmt.Println(fv.Int())// 2

	// Try to set the value of len.
	fv.Set(reflect.ValueOf(3)) // ILLEGAL
}