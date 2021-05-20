package main

import (
	"fmt"
	"reflect"
)

func main() {

	// tweecking index position with custom positions
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
		11: "November",``
		45: "December"}
	fmt.Println(months)
	fmt.Println(reflect.TypeOf(months).Kind())

	fmt.Println("months[3]=", months[3])
	fmt.Println("months[12]=", months[12])
	fmt.Println("months[45]=", months[45])

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
