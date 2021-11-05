package main

import "fmt"

func main() {
	type T struct {
		a int
		b float64
		c string
	}

	t := &T{7, -2.35, "abc\tdef"}

	fmt.Printf("%v\n", t)  // &{7 -2.35 abc   def}
	fmt.Printf("%+v\n", t) // &{a:7 b:-2.35 c:abc     def}
	fmt.Printf("%#v\n", t) // &main.T{a:7, b:-2.35, c:"abc\tdef"}
	fmt.Printf("%T\n", t)  // *main.T

}
 