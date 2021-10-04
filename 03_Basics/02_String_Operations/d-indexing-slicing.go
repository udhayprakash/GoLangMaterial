package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(`len("Hello world")	=`, len("Hello world"))

	name := "Hello world"
	fmt.Println("name				=", name)
	fmt.Println("len(name)			=", len(name))
	fmt.Println("utf8.RuneCountInString(name)=", utf8.RuneCountInString(name))

	// Indexing - starts with zero
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10
	fmt.Println()
	fmt.Println("name[0]			=", name[0])
	fmt.Printf("name[0]			= %c\n", name[0]) // ascii value
	fmt.Printf("name[4]			= %c\n", name[4])
	fmt.Printf("name[5]			= %q\n", name[5])
	fmt.Printf("name[10]		= %q\n", name[10])
	// fmt.Printf("name[11]		= %q\n", name[11])
	// panic: runtime error: index out of range

	fmt.Printf("name[-0]		= %c\n", name[-0])
	// fmt.Printf("name[-1]		= %c\n", name[-1])
	// invalid string index -1 (index must be non-negative)
	fmt.Println()

	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10

	// slicing - last position char is not included
	fmt.Println("name[2:7]			=", name[2:7])
	fmt.Println("name[1:10]			=", name[1:10])
	fmt.Println("name[1:11]			=", name[1:11])
	// fmt.Println("name[1:12]			=", name[1:12])
	// slice bounds out of range

	fmt.Println()
	fmt.Println("name[1:]			=", name[1:]) // default final index -> len(string)
	fmt.Println("name[:]			=", name[:])   // default start index -> 0

	// fmt.Println("name[1:11:1]			=", name[1:11:1]) // invalid operation name[1:11:1] (3-index slice of string)

	fmt.Println()
	fmt.Println("name[len(name)- 3:]=", name[len(name)-3:])
	fmt.Println("name[:len(name)- 3]=", name[:len(name)-3])

	// To print 'world'
	fmt.Println(name[6:11], name[len(name)-5:])
	fmt.Println()

	// Capitalize first half of the string
	s := "i love food"
	half := len(s) / 2
	fmt.Println("s    = ", s)
	fmt.Println("s[:half]=", s[:half])
	result := strings.ToUpper(s[:half]) + strings.ToLower(s[half:])
	fmt.Println(result)
}
