package main

import (
	"fmt"
	"reflect"
)

type fruit string

func main()  {
	var apple fruit = "Kashmir apple"
	fmt.Println(apple, reflect.TypeOf(apple))
}