package main

import (
	"fmt"
	"reflect"
)
func main(){
	a1 := [...]string{"a", "b", "c", "d", "e", "f", "g"} // len-7; cap-7
	//                 0    1    2    3    4    5    6
	// last index is not included, in slice
	s1 := a1[2:4] // [c d] 				len-2; cap-5
	s2 := a1[3:6] // [d e f] 			len-3; cap-4
	s3 := a1[1:5] // [b c d e] 			len-4; cap-6
	s4 := a1[:]   // [a b c d e f g] 	len-7; cap-7
	s5 := a1[:7]  // [a b c d e f g] 	len-7; cap-7
	// s6 := a1[:8]  - error - invalid slice index 8 (out of bounds for 7-element array)

	fmt.Println("\na1=", a1, reflect.TypeOf(a1).Kind())
	fmt.Println("len(a1)=", len(a1))
	fmt.Println("cap(a1)=", cap(a1))

	fmt.Println("\ns1=", s1, reflect.TypeOf(s1).Kind())
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))

	fmt.Println("\ns2=", s2, reflect.TypeOf(s2).Kind())
	fmt.Println("len(s2)=", len(s2))
	fmt.Println("cap(s2)=", cap(s2))

	fmt.Println("\ns3=", s3, reflect.TypeOf(s3).Kind())
	fmt.Println("len(s3)=", len(s3))
	fmt.Println("cap(s3)=", cap(s3))

	fmt.Println("\ns4=", s4, reflect.TypeOf(s4).Kind())
	fmt.Println("len(s4)=", len(s4))
	fmt.Println("cap(s4)=", cap(s4))

	fmt.Println("\ns5=", s5, reflect.TypeOf(s5).Kind())
	fmt.Println("len(s5)=", len(s5))
	fmt.Println("cap(s5)=", cap(s5))

}


