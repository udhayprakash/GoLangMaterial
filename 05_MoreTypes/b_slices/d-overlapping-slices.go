package main

import (
	"fmt"
	"reflect"
)

func main() {

	// tweecking index position with custom positions
	months := [...]string{

		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",

		1: "January",
		2: "February",
		3: "March",
		4: "April",

		45: "December"}
	fmt.Println(months)
	fmt.Println(len(months))                   // 46
	fmt.Println(reflect.TypeOf(months).Kind()) // array

	fmt.Println("months[3]=", months[3]) // March
	fmt.Println("months[12]=", months[12])
	fmt.Println("months[45]=", months[45]) // December

	Q2 := months[4:7]
	fmt.Println("Q2    :", Q2) // ["April" "May" "June"]
	for _, q := range Q2 {
		fmt.Println("q = ", q)
	}
	fmt.Println()

	summer := months[6:9]
	fmt.Println("summer:", summer) // ["June" "July" "August"]

	for _, s := range summer {
		fmt.Println("s =", s)
	}

	fmt.Println()

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Println("s =", s, " q = ", q)
			}
		}
	}
}
