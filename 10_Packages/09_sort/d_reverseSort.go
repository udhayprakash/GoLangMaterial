package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Example For Integer Reverse Sort")
	num := []int{70, 80, 20, 50, 10}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println("\t", num)

	fmt.Println("Example For String Reverse Sort")
	str := []string{"GFG", "Rank", "India", "Amid", "Covid19"}
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println("\t", str)
}
