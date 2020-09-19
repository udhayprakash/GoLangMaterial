package main

import (
	"bytes"
	"fmt"
)

func main(){
	b1 := byte('a')
	b2 := []byte("A")
	b3 := []byte{'a', 'b', 'c'}

	fmt.Printf("b1 = %c\n", b1) // a
	fmt.Printf("b2 = %c\n", b2) // [A]
	fmt.Println()

	fmt.Printf("b3 = %c\n", b3) // [a b c]
	fmt.Printf("b3 = %s\n", b3) // abc
	fmt.Println()

	s1 := []byte("Hello")
	s2 := []byte("World")
	s3 := [][]byte{s1, s2}
	s4 := bytes.Join(s3, []byte(","))
	s5 := []byte{}
	s5 = bytes.Join(s3, []byte("--"))
	s6 := [][]byte{[]byte("one"), []byte("two"), []byte("three")}

	fmt.Printf("s1 = %s\n", s1) // Hello
	fmt.Printf("s2 = %s\n", s2) // World
	fmt.Printf("s3 = %s\n", s3) // [Hello World]
	fmt.Printf("s4 = %s\n", s4) // Hello,World
	fmt.Printf("s5 = %s\n", s5) // Hello--World
	fmt.Printf("s6 = %s\n", s6) // [one two three]
	fmt.Printf("%s\n", bytes.Join(s6, []byte(", ")))
	// one, two, three



}