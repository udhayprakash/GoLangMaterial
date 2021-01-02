package main

import (
	"fmt"
	"math/big"
)

func main() {

	// base 10
	bigInteger := new(big.Int)
	bigInteger.SetString("100000000000000000L", 10) //18

	fmt.Println("Value : ", bigInteger)
	fmt.Println("Length : ", len(bigInteger.Text(10)))
}
