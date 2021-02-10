package main

import "fmt"

func main() {
	vals := make(chan float64, 4)

	vals <- 85.4
	vals <- 0.6
	vals <- -54.9
	vals <- 32.14

	if vals == vals {
		fmt.Println(3)
		if vals == vals {
			fmt.Println(0)
			if <-vals == <-vals {
				fmt.Println(5)
			} else {
				if <-vals == <-vals {
					fmt.Println(8)
				}
			}
		}
	}
}
