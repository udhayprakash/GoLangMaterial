package main 

import "fmt"

func main() {
	a := [1, 3, 5, 7, 11]
	b := [x*x for x <- a, x > 3]
	println(b)
}

// NOTE: https://github.com/qiniu/goplus