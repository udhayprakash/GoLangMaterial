package main

import (
	"fmt"
	"strconv"
)

func gt(arg0 string, arg1 string) int {
	var1, _ := strconv.Atoi(arg0)
	var2, _ := strconv.Atoi(arg1)

	if var1 > var2 {
		return var1
	}
	return var2
}

func main() {
	o := gt("50", "10")
	fmt.Println(o)
}
