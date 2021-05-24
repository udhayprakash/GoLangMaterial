package main

import (
	"errors"
	"fmt"
)

func CirlceArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

func main() {
	fmt.Println(CirlceArea(10))  // 100 <nil>
	fmt.Println(CirlceArea(-10)) // 0 Provide Positive NUmber

	area, err := CirlceArea(-1)
	if err != nil {
		fmt.Println("Error Encountered..., ", err)
		return
	}
	fmt.Println("area =", area)
}
