package main

import "fmt"

func main() {
	fmt.Println(`len("Hello world")	=`, len("Hello world"))

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
	//fmt.Println("name[11]			=", name[11]) // panic: runtime error: index out of range [11] with length 11

	fmt.Println("name[-0]			=", name[-0])
	//fmt.Println("name[-1]			=", name[-1]) // invalid string index -1 (index must be non-negative)

	// slicing - last position char is not included
	fmt.Println("name[1:6]			=", name[1:6])
	fmt.Println("name[1:10]			=", name[1:10])
	fmt.Println("name[1:11]			=", name[1:11])
	//fmt.Println("name[1:12]			=", name[1:12])
	//panic: runtime error: slice bounds out of range [:12] with length 11

	fmt.Println()
	fmt.Println("name[1:]			=", name[1:]) // default final index - len(string)
	fmt.Println("name[:]			=", name[:])   // default start index - 0

	//fmt.Println("name[1:11:1]			=", name[1:11:1]) // invalid operation name[1:11:1] (3-index slice of string)

	fmt.Println()
	fmt.Println("name[len(name)- 3:]=", name[len(name)-3:])
	fmt.Println("name[:len(name)- 3]=", name[:len(name)-3])

}
