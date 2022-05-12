package main

import (
	"fmt"
	"reflect"
)

/*
Purpose:
	To allocate memory Go has two primitives, new and make.

	new allocates; make initializes

	new(T)
		- returns *T pointing to a zeroed T
		- allocates zeroed storage for a new item of type T
			and returns its address, a value of type *T.
		- it returns a pointer to a newly allocated zero
			value of type T. This is important to remember.

	make(T, args)
		- returns an initialized (not zero!) value of type T,
			and not a pointer: *T.
		- It creates slices, maps, and channels only
*/

func main() {
	newVar := new([]int)
	makeVar := make([]int, 5)
	fmt.Println(newVar, reflect.TypeOf(newVar).Kind())   // &[] ptr
	fmt.Println(makeVar, reflect.TypeOf(makeVar).Kind()) // [0 0 0 0 0]  slice

	fmt.Println(new([]int))     // &[]
	fmt.Println(make([]int, 5)) // [0 0 0 0 0]

	fmt.Println(new([]float64))     // &[]
	fmt.Println(make([]float64, 5)) // [0 0 0 0 0]

	fmt.Println(new([]string))     // &[]
	fmt.Println(make([]string, 5)) // [     ]

	fmt.Println(new([]rune))     // &[]
	fmt.Println(make([]rune, 5)) // [0 0 0 0 0]

	fmt.Println(new([]bool))     // &[]
	fmt.Println(make([]bool, 5)) // [false false false false false]

	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})

	fmt.Println("ci=", ci)
	fmt.Println("cs=", cs)
	fmt.Println("cf=", cf)

}
