package main

import (
	"fmt"
)

type NegativeRadiusError struct {
	radius int
}

func (e *NegativeRadiusError) Error() string {
	return fmt.Sprintf("Invalid radius %v for calculating area", e.radius)
}

func CircleArea(radius int) (int, error) {
	if radius < 0 {
		// return 0, errors.New("Provide Positive Number")
		return 0, &NegativeRadiusError{radius: radius}
	}
	return radius * radius, nil
}

func main() {
	fmt.Println(CircleArea(10)) // 100 <nil>
	fmt.Println(CircleArea(5))  // 25 <nil>
	fmt.Println(CircleArea(0))  // 0 <nil>
	fmt.Println(CircleArea(-5)) // 0 Provide Positive Number

	area, err := CircleArea(-1)
	if err != nil {
		fmt.Println("Error Encountered..., ", err)
		return
	}
	fmt.Println("area =", area)
}
