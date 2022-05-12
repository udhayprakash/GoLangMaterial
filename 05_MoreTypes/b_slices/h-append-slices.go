package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s) // [  ]

	s[0] = "a"
	fmt.Println("s =", s) // [a  ]

	s[1], s[2] = "b", "c" // tuple unpacking
	fmt.Println("s =", s) // [a b c]

	fmt.Println("len:", len(s)) // 3
	fmt.Println("cap:", cap(s)) // 3

	// s[3] = "d"
	// panic: runtime error: index out of range [3] with length 3

	// To add elements beyond slice length,
	s = append(s, "d")
	fmt.Println("s =", s) //  [a b c d]

	s = append(s, "e", "f", "g", "h")
	fmt.Println("s =", s)       // [a b c d e f g h]
	fmt.Println("len:", len(s)) // 8
	fmt.Println()

	// creating slice with len & capacity limits
	s2 := make([]string, 3, 5)
	fmt.Println("s2 =", s2)           // [  ]
	fmt.Println("len(s2) =", len(s2)) // 3
	fmt.Println("cap(s2) =", cap(s2)) // 5

	s2[0], s2[1], s2[2] = "z", "y", "x"
	fmt.Println("s2 =", s2) // [z y z]

	// s2[3], s2[4] = "w", "v"
	// panic: runtime error: index out of range [3] with length 3

	s2 = append(s2, "w", "v")
	fmt.Println("s2 =", s2)           // [z y x w v]
	fmt.Println("len(s2) =", len(s2)) // 5
	fmt.Println("cap(s2) =", cap(s2)) // 5

	// NOTE: added elements beyond slice caapcity will recreate the slice with new cap/len
	s2 = append(s2, "u")
	fmt.Println("s2 =", s2)           //  [z y x w v u]
	fmt.Println("len(s2) =", len(s2)) // 6
	fmt.Println("cap(s2) =", cap(s2)) // 10
	fmt.Println()

	//==============================================
	// Dynamic slice
	c := make([]string, len(s2))
	fmt.Println("c =", c) //  [     ]

	copy(c, s2)                       // safe copy
	fmt.Println("cpy:", c)            //[z y x w v u]
	fmt.Println("len(s2) =", len(s2)) // 6
	fmt.Println("cap(s2) =", cap(s2)) // 10

	// In this case, changes in one slice, were not leaked into another slice
	c[3] = "ww"
	fmt.Println("c =", c)  //  [z y x ww v u]
	fmt.Println("s2=", s2) // [z y x w v u]  -- NOT EFFECTED

}
