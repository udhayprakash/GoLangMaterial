package main

import "fmt"

/*
Pointers to struct
==================
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct
pointer p we could write (*p).X. However, that notation is
cumbersome, so the language permits us instead to write
just p.X, without the explicit dereference.
*/

type myVertex struct {
	X int
	Y int
}

func main(){
	value := myVertex{1, 2}
	valuePtr := &value

	value.X = 222
	value.Y = 333

	valuePtr.X = 444
	valuePtr.Y = 555

	fmt.Println("value.X == valuePtr.X :", value.X == valuePtr.X)
}
