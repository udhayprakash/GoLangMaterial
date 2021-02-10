package main

import (
	"fmt"
	"reflect"
)

type keet struct {
	F string `species:"gopher" color:"blue"`
	G string `species:"hopper" color:"green"`
}

func main() {
	k := keet{}
	val := reflect.TypeOf(k)
	next := val.Field(1)

	fmt.Println(next.Tag.Get("color"), next.Tag.Get("species"))
}
