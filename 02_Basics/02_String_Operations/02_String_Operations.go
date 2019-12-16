package main

import "fmt"

func main() {
	fmt.Println(`len("Hello world") =`, len("Hello world"))

	name := "Hello world"
	fmt.Println("name				=", name)
	fmt.Println("len(name)			=", len(name))


	// Indexing - starts with zero
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10
	fmt.Println()
	fmt.Println("name[0]			=", name[0])
	fmt.Println("name[1]			=", name[1])
	fmt.Println("name[5]			=", name[5])
	fmt.Println("name[10]			=", name[10])
	// fmt.Println("name[11]			=", name[11])

	fmt.Println("name[-0]			=", name[-0])
	// fmt.Println("name[-1]			=", name[-1])

	// slicing - last position char is not included
	fmt.Println("name[1:6]			=", name[1:6])
	fmt.Println("name[1:10]			=", name[1:10])
	fmt.Println("name[1:11]			=", name[1:11])
	// fmt.Println("name[1:45]			=", name[1:45])
	fmt.Println("name[1:]			=", name[1:])
	fmt.Println("name[:]			=", name[:])

	// fmt.Println("name[1:11:1]			=", name[1:11:1])

	fmt.Println("name[8:]			=", name[8:])
	fmt.Println("name[len(name)- 3:]=", name[len(name)- 3:])


}