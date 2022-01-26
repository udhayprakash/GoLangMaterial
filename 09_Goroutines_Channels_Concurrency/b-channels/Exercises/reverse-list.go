package main

import (
	"fmt"
)

func reverse(lst []string) chan string {
	ret := make(chan string)
	go func() {
		for i := range lst {
			ret <- lst[len(lst)-1-i]
		}
		close(ret)
	}()
	return ret
}

func main() {
	elms := []string{"a", "b", "c", "d"}
	for e := range reverse(elms) {
		fmt.Println(e)
	}
}
