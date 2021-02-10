package main

import (
	"fmt"
	"reflect"
)

type Deviant struct {
	Art *Deviant
}

func main() {
	p := Deviant{Art: &Deviant{}}
	q := Deviant{Art: &Deviant{}}

	fmt.Println(reflect.DeepEqual(p, q)) // true

}
