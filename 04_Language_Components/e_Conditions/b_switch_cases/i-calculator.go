package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	default:
		fmt.Println("Invalid operation")
	}
	fmt.Println("result = ", result)
}

func getInput(prompt string) float64 {
	fmt.Printf("%v: ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}
func getOperation() string {
	fmt.Print("Select an operation (+, -, *, /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValues(v1, v2 float64) float64 {
	return v1 + v2
}

func subtractValues(s1, s2 float64) float64 {
	return s1 - s2
}


// Assignment 1 - complete the code, by adding functionality for  *, /
// 						condition for denominator as 0, when division operation

// Assignment 2 - deisgn a cacluclator which will take the entire expression in one line
//			100 + 200
