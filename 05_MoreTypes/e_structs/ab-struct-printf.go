package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		a int
		b float64
		c string
	}

	t1 := T{a: 10, b: 20.2, c: "30"}

	fmt.Printf("t1 =%v\n", t1)             // {10 20.2 30}
	fmt.Printf("t1 =%+v\n", t1)            // {a:10 b:20.2 c:30}
	fmt.Printf("t1 =%#v\n", t1)            // main.T{a:10, b:20.2, c:"30"}
	fmt.Printf("t1 =%T\n", t1)             // main.T
	fmt.Println(reflect.TypeOf(t1).Kind()) // struct
	fmt.Println()

	t2 := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("t2 =%v\n", t2)             // &{7 -2.35 abc       def}
	fmt.Printf("t2 =%+v\n", t2)            // &{a:7 b:-2.35 c:abc def}
	fmt.Printf("t2 =%#v\n", t2)            // &main.T{a:7, b:-2.35, c:"abc\tdef"}
	fmt.Printf("t2 =%T\n", t2)             // *main.T
	fmt.Println(reflect.TypeOf(t2).Kind()) // prt

}
