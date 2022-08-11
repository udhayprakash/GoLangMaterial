package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s:= []string{"Black Plum", "Red Plum", "Orange", "Kiwi", "Mango"}
	indexes := rand.Perm(len(s))
	fmt.Println(indexes) // [0 4 2 3 1]

	for i := 0; i < len(indexes); i++ {
		fmt.Println(s[indexes[i]])
	}
}