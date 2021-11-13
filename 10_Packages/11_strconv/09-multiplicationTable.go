package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("os.Args[1]=%v Type=%[1]T\n", os.Args[1])

	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("val=%v Type=%[1]T\n", val)
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", val, i, val*i)
	}
}
