package main

import "fmt"

func main() {
	vals := make(chan float64, 4)

	vals <- 85.4
	vals <- 0.6
	vals <- -54.9
	vals <- 32.14

	if vals == vals { // Channel Comparison - channel compared to itself
		fmt.Println(3)
		if vals == vals { // Channel Comparison - channel compared to itself
			fmt.Println(0)
			if <-vals == <-vals { // compares first value with second, from channel
				fmt.Println(5)
			} else {
				if <-vals == <-vals { // compares third, with fourth
					fmt.Println(8)
				}
			}
		}
	}
}
