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

	type Even bool
	const (
		// 0 % 2 == 0 ==> Even(true)
		a = Even(iota % 2 == 0)
		// 1 % 2 == 0 ==> Even(false)
		b
		// 2 % 2 == 0 ==> Even(true)
		c
		// 3 % 2 == 0 ==> Even(false)
		d
	)
	fmt.Print("a, b, c, d : ")
	fmt.Println(a, b, c, d)
}