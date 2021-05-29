package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}


func main() {
	nextEven := makeEvenGenerator()
	fmt.Println("nextEven() =", nextEven()) // 2
	fmt.Println("nextEven() =", nextEven()) // 4
	fmt.Println("nextEven() =", nextEven()) // 6

}
