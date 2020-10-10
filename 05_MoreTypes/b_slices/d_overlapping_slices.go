package main

import "fmt"

func main() {

	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}

	Q2 := months[4:7]
	fmt.Println("Q2    :", Q2) // ["April" "May" "June"]
	fmt.Println()

	summer := months[6:9]
	fmt.Println("summer:", summer) // ["June" "July" "August"]

	for _, s := range summer {
		fmt.Println("s =", s)
	}
	for _, q := range Q2 {
		fmt.Println("q = ", q)
	}
	fmt.Println()

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Println("s =", s)
				fmt.Println("q = ", q)
			}
		}
	}

}
