package main

import "fmt"

func length(args ...int) int {
	return len(args)
}

func main() {
	vs := []int{24, 55, 88, 99, 4, 2}
	fmt.Println(length(vs...))
}


// -race is the flg need to be added in command line 
// to find data race