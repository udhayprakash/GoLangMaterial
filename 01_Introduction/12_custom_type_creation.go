package main

import (
	"fmt"
	"reflect"
)

type fruit string

func main()  {
	var apple fruit = "Kashmir apple"
	fmt.Println(apple, reflect.TypeOf(apple)) // Kashmir apple main.fruit

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Wednesday, reflect.TypeOf(Wednesday)) // 3 main.Weekday
}