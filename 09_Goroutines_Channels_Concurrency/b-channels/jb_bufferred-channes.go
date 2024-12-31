package main

import "fmt"

func main() {
	vals := make(chan float64, 4)

	vals <- 11.1
	vals <- 11.1 //22.2
	vals <- 33.3
	vals <- 44.4

	if vals == vals { // Channel Comparison - channel compared to itself
		fmt.Println("condition 1 passed", vals)
		if vals == vals { // Channel Comparison - channel compared to itself
			fmt.Println("condition 2 passed", vals)
			if <-vals == <-vals { // compares first value with second, from channel
				fmt.Println("First two values are equal")

			} else {
				if <-vals == <-vals {
					fmt.Println("third and fourth values are equal")
				}
			}
		}
	}
}
