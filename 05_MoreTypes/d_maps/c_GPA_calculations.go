package main

import "fmt"

var POINTS = map[string]float32{
	"A+": 4.0,
	"A":  4.0,
	"A-": 3.67,
	"B+": 3.33,
	"B":  3.0,
	"B-": 2.67,
	"C+": 2.33,
	"C":  2.0,
	"C-": 1.67,
	"D+": 1.33,
	"D":  1.0,
	"F":  0.0,
}

func main() {
	fmt.Println("Enter all your letter grades, one per line")
	fmt.Println("Enter a blank line to designate the end")
	fmt.Println(POINTS)

	var grade string

	for {
		fmt.Scanf("%s", &grade)
		fmt.Println("grade: ", grade, "\tpoints: ", POINTS[grade])
	}
}
