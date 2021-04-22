package main

import (
	"fmt"
)

/*
Write a function that uses the minimum amount
of coins necessary to full fill the order.
The coins are 25c, 15c, 10c, 1c.
*/

func main() {
	fmt.Println("123 =>", GetMinimumCoins(123)) // map[1:8 15:1 25:4]
	fmt.Println("400 =>", GetMinimumCoins(400)) // map[25:16]
	fmt.Println("111 =>", GetMinimumCoins(111)) // map[1:1 10:1 25:4]
	fmt.Println("1   =>", GetMinimumCoins(1))   // map[1:1]

}

func GetMinimumCoins(amount int) map[int]int {
	coins := []int{25, 15, 10, 1}
	coinsCount := make(map[int]int)
	for _, coin := range coins {
		if (amount / coin) != 0 {
			coinsCount[coin] = (amount / coin)
		}
		amount -= (amount / coin) * coin
		if amount == 0 {
			break
		}
	}
	return coinsCount
}
