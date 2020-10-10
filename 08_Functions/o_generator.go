package main

import "fmt"

//func makeEvenGenerator() func() uint {
//	i := uint(0)
//	return func() uint {
//		i += 2
//		return i
//	}
//}

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
	fmt.Println("nextEven() =", nextEven())
	fmt.Println("nextEven() =", nextEven())
	fmt.Println("nextEven() =", nextEven())
}
