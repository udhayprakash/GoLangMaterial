package main

import "fmt"

func main() {
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			fmt.Printf("A\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	A	A	A	A
	A	A	A	A	A
	A	A	A	A	A
	A	A	A	A	A
	A	A	A	A	A
	*/
	fmt.Println("\n\n====================")

	for _, alphabet := range "ABCDE" {
		for i = 0; i < 5; i++ {
			fmt.Print(string(alphabet), "\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	A	A	A	A
	B	B	B	B	B
	C	C	C	C	C
	D	D	D	D	D
	E	E	E	E	E
	*/

	fmt.Println("\n\n====================")

	for i = 0; i < 5; i++ {
		for _, alphabet := range "ABCDE" {
			fmt.Print(string(alphabet), "\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	B	C	D	E
	A	B	C	D	E
	A	B	C	D	E
	A	B	C	D	E
	A	B	C	D	E
	*/
	fmt.Println("\n\n====================")

	for i = 65; i < 90; i++ {
		if (i != 65) && (i%5 == 0) {
			fmt.Println()
		}
		fmt.Print(string(i), "\t")
	}
	/*OUTPUT:
	A	B	C	D	E
	F	G	H	I	J
	K	L	M	N	O
	P	Q	R	S	T
	U	V	W	X	Y
	*/
	fmt.Println("\n\n====================")
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			fmt.Print(string(i+j+65), "\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	B	C	D	E
	B	C	D	E	F
	C	D	E	F	G
	D	E	F	G	H
	E	F	G	H	I
	*/
	fmt.Println("\n\n====================")
	for j = 0; j < 5; j++ {
		for i = 69; i > 64; i-- {
			fmt.Print(string(i), "\t")
		}
		fmt.Println()
	}

	/*OUTPUT:
	E	D	C	B	A
	E	D	C	B	A
	E	D	C	B	A
	E	D	C	B	A
	E	D	C	B	A
	*/

	fmt.Println("\n\n====================")
	for i = 69; i > 64; i-- {
		for j = 0; j < 5; j++ {
			fmt.Print(string(i), "\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	E	E	E	E	E
	D	D	D	D	D
	C	C	C	C	C
	B	B	B	B	B
	A	A	A	A	A
	*/
	fmt.Println("\n\n====================")
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if i == 0 || i == 4 || j == 0 || j == 4 {
				fmt.Print("A\t")
			} else {
				fmt.Print(" \t")
			}
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	A	A	A	A
	A	 	 	 	A
	A	 	 	 	A
	A	 	 	 	A
	A	A	A	A	A
	*/
	fmt.Println("\n\n====================")
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if i >= j {
				fmt.Print("A\t")
			} else {
				fmt.Print(" \t")
			}
		}
		fmt.Println()
	}
	/*OUTPUT:
	A
	A	A
	A	A	A
	A	A	A	A
	A	A	A	A	A
	*/
	fmt.Println("\n\n====================")
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if i <= j {
				fmt.Print("A\t")
			} else {
				fmt.Print(" \t")
			}
		}
		fmt.Println()
	}
	/*OUTPUT:
	A	A	A	A	A
	 	A	A	A	A
	 	 	A	A	A
	 	 	 	A	A
	 	 	 	 	A
	*/
	fmt.Println("\n\n====================")
	for i = 0; i < 5; i++ {
		for _, eachChar := range "Prakash"{
			fmt.Print(string(eachChar), "\t")
		}
		fmt.Println()
	}
	/*OUTPUT:
	P	r	a	k	a	s	h
	P	r	a	k	a	s	h
	P	r	a	k	a	s	h
	P	r	a	k	a	s	h
	P	r	a	k	a	s	h
	*/

}
