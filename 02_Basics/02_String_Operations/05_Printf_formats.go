package main

import "fmt"

func main() {
	s := "Hello world"

	fmt.Println("len(s) =", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	// for i:=0; i < len(s); i++ {
	// 	fmt.Printf("%v", s[i])
	// }
	// fmt.Println()

	// for i:=0; i < len(s); i++ {
	// 	fmt.Printf("%x", s[i])
	// }
	// fmt.Println()

	// for i:=0; i < len(s); i++ {
	// 	fmt.Printf("%T", s[i])
	// }
	// fmt.Println()

}
