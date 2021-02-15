/*
--------------------------------------------------------------------------------
From	b []byte	i []int		r []rune	s string	f float32	i int
To
--------------------------------------------------------------------------------
[]byte	·									[]byte(s)
[]int				·						[]int(s)
[]rune										[]rune(s)
string	string(b)	string(i)	string(r)	·
float32													·			float32(i)
int														int(f)		·
--------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	// string to a slice of bytes or runes
	mystring := "hello this is string"
	fmt.Println("mystring  =", mystring)

	byteslice := []byte(mystring)
	fmt.Println("byteslice =", byteslice)

	runeslice := []rune(mystring)
	fmt.Println("runeslice =", runeslice)

	//  slice of bytes or runes to a string
	b := []byte{'h', 'e', 'l', 'l', 'o'} // Composite literal.
	fmt.Println("\nb=", b)

	s := string(b)
	fmt.Println("s=", s)

	i := []rune{257, 1024, 65}
	fmt.Println("i=", i)

	r := string(i)
	fmt.Println("r=", r)

}
