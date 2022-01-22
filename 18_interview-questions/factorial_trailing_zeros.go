package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func trailingZeros(num int) int {
	lastdigit, noOfZeros := 0, 0
	for {
		lastdigit = num % 10
		if lastdigit != 0 {
			break
		}
		noOfZeros++
		num = num / 10
	}
	return noOfZeros
}

func main() {
	fmt.Println(trailingZeros(factorial(10)) == 2)
	for i := 0; i < 20; i++ {
		fmt.Println(i, trailingZeros(factorial(i)), factorial(i))
	}
	fmt.Println(factorialBig(big.NewInt(103)))

}

func factorialBig(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorialBig(n.Sub(x, n)))
}
