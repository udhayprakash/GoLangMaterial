package main

import "fmt"

/*
Looping over elements in slices, arrays, maps, channels
and strings is often better done using the range keyword:
*/
func main() {
	strings := []string{"hello", "world", "Go", "Language"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
