package main

import (
	"fmt"
	"reflect"
)
/*
NOTE: 
key type must have an == operation (NO  maps, slices or funcs as keys)
operations run in constant expected time O(1)

Construct		m := map[key]value{}
Insert 			m[k] = v
LookUp			v, isKeyPresent = m[k]
Delete			delete(m, k)
Iterate			for k, v := range m
Size			len(m)
*/
func main() {
	var alphabets = map[string]string{
		"b": "banana",
		"a": "apple",
		"c": "cat", // last comma is mandatory
	}
	fmt.Printf("alphabets=%v - %[1]T\n", alphabets) // [a:apple b:banana c:cat]
	fmt.Println(reflect.TypeOf(alphabets), reflect.TypeOf(alphabets).Kind())

	another := map[bool]string{
		true:  "This is true",
		false: "This is false",
	}
	fmt.Println("another=", another)
	// Its rare to use bool as key

	someThing := map[int]bool{
		123:  true,
		-234: false,
		34:   true,
		5:    true,
		// 123: false, // duplicate key 123 in map literal
	}
	fmt.Println("someThing = ", someThing)          // map[-234:false 5:true 34:true 123:true]
	fmt.Println("len(someThing) =", len(someThing)) // 4
	// fmt.Println("cap(someThing) =", cap(someThing)) // doesnt work

	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	// For maps, Printf sorts the output in lexicographically by key.
	fmt.Printf("%#v\n", timeZone)
}

// NOTE: map stores the value in ascending order of the keys