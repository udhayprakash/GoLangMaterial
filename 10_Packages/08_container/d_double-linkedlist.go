package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	var x list.List
	fmt.Println("x                       :", x)
	fmt.Println("reflect.TypeOf(x)       :", reflect.TypeOf(x))
	fmt.Println("reflect.TypeOf(x).Kind():", reflect.TypeOf(x).Kind())

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	fmt.Println("x                       :", x)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
