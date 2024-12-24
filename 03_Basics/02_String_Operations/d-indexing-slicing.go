package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(`len("Hello world")	=`, len("Hello world")) // 11

	name := "Hello world"
	fmt.Println("name				=", name)
	fmt.Println("len(name)			=", len(name))
	fmt.Println("utf8.RuneCountInString(name)=", utf8.RuneCountInString(name))
	fmt.Println()

	name2 := "హలో వరల్డ్"
	fmt.Println("name2				=", name2)
	fmt.Println("len(name2)			=", len(name2))
	fmt.Println("utf8.RuneCountInString(name2)=", utf8.RuneCountInString(name2))
	fmt.Println()

	// Indexing - starts with zero
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10
	fmt.Println("name				=", name)
	fmt.Println("name[0]			=", name[0])
	fmt.Printf("name[0]			= %c\n", name[0]) // ascii value
	fmt.Printf("name[0]		    = %q\n", name[0])

	fmt.Printf("name[5]			= %q\n", name[5])
	fmt.Printf("name[10]		= %q\n", name[10])

	// fmt.Printf("name[11]		= %q\n", name[11])
	// panic: runtime error: index out of range

	fmt.Printf("name[-0]		= %q\n", name[-0])
	// fmt.Printf("name[-2]		= %q\n", name[-2])
	//  index -2 (constant of type int) must not be negative

	// ====== SLICING
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10

	// string[start_index: end_index]
	// slicing - last position char is not included
	fmt.Printf("name[2]			=%q\n", name[2])
	fmt.Printf("name[7]			=%q\n", name[7])

	fmt.Println("name[2:7]			=", name[2:7])   //  llo w
	fmt.Println("name[1:10]			=", name[1:10]) // ello	worl
	fmt.Println("name[1:11]			=", name[1:11]) // ello	world

	// fmt.Println("name[1:12]			=", name[1:12])
	// panic: runtime error: slice bounds out of range [:12] with length 11

	/*
		default
			start index = 0
			final index = len(string)
	*/
	fmt.Println("name[1:]			=", name[1:]) // default final index -> len(string)
	fmt.Println("name[:]			=", name[:])   // default start index -> 0

	// fmt.Println("name[1:11:1]			=", name[1:11:1])
	// error: invalid 3-index slice of string

	fmt.Println()
	// H	e	l	l	o	 	w	o	r	l	d
	// 0   	1	2	3	4	5	6	7	8	9	10

	// To Get last N characters;  len(name)= 11
	// 11 - 3 = 8  => name[8:]
	fmt.Println("name[8:]             =", name[8:])           // rld
	fmt.Println("name[len(name) - 3:] =", name[len(name)-3:]) // rld

	fmt.Println("name[len(name)- 5:]=", name[len(name)-5:]) // world

	// To Get from start till last N characters;  len(name)= 11
	fmt.Println("name[:8]             =", name[:8])           // Hello wo
	fmt.Println("name[:len(name) - 3] =", name[:len(name)-3]) // Hello wo

	fmt.Println("name[:len(name) - 5] =", name[:len(name)-5]) // Hello
	fmt.Println()

	// Capitalize first half of the string
	sentence := " I love eating mangos in summer"

	halfLen := len(sentence) / 2
	fmt.Printf("sentence            = %s\n", sentence)
	fmt.Printf("sentence[:halfLen]  = %s\n", sentence[:halfLen])
	fmt.Printf("sentence[halfLen:]  = %s\n", sentence[halfLen:])

	res := strings.ToUpper(sentence[:halfLen]) + strings.ToLower(sentence[halfLen:])
	fmt.Println("result = ", res)

}
