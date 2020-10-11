package main

import "fmt"

/*
_Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
(sometimes called _hashes_ or _dicts_ in other languages).

Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).

map keys may be of any type that is comparable.
comparable types: boolean, numeric, string, pointer, channel, and interface types,
	   and structs or arrays that contain only those types.

- unordered collection of key-value pairs
	- keys are any type that supports == and != operators
	- values are any type

*/
func main(){
	m := make(map[string]int)

	m["suresh"] = 100
	m["Naresh"] = 200
	fmt.Println("map         :", m)  // map[Naresh:200 suresh:100]
	fmt.Println("len         :", len(m))	// 2


	m["Ganesh"] = 300
	m["Mahesh"] = 400
	fmt.Printf("map         : %v\n", m) //map[Ganesh:300 Mahesh:400 Naresh:200 suresh:100]
	fmt.Println("len         :", len(m))	// 4

	delete(m, "Mahesh")
	fmt.Println("map         :", m) //map[Ganesh:300 Naresh:200 suresh:100]


	val, isKeyPresent := m["Naresh"]
	fmt.Println("\nval         :", val)
	fmt.Println("isKeyPresent:", isKeyPresent)


	val1 := m["Ganesh"]
	fmt.Println("\nval1        :", val1) // 300

	//blank identifier _ -used when we dont need a value
	_, isKeyPresent = m["superman"]
	fmt.Println("isKeyPresent:", isKeyPresent) // false






}