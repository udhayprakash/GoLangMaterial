package main

import "fmt"

func OddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextOdd := OddGenerator()
	fmt.Println("nextOdd() =", nextOdd())
	fmt.Println("nextOdd() =", nextOdd())
	fmt.Println("nextOdd() =", nextOdd())
}
