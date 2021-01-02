package main

import (
	"fmt"
)

func PercentageChange(old, new int) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}

func main() {

	fmt.Printf("Old : 500, New : 500 Change : %0.2f %% \n", PercentageChange(500, 500))

	// anything divided by 0 will become infinity
	fmt.Printf("Old : 0, New : 50 Change : %0.2f %% \n", PercentageChange(0, 50))

	fmt.Printf("Old : 100, New : 30 Change : %0.2f %% \n", PercentageChange(100, 30))

	fmt.Printf("Old : 1000, New : 830 Change : %0.2f %% \n", PercentageChange(1000, 830))

	fmt.Printf("Old : 5, New : 3 Change : %0.2f %% \n", PercentageChange(5, 3))

	fmt.Printf("Old : -15, New : 3 Change : %0.2f %% \n", PercentageChange(-15, 3))

	fmt.Printf("Old : 100, New : 130 Change : %0.2f %% \n", PercentageChange(100, 130))
}
