// Program to display the command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println(`
			USAGE  : 1_dinnertotal <Total Meal Amount> <Tip Percentage>
			EXAMPLE: 1_dinnertotal 2000 10
			`)
	}
	if len(args) != 2 {
		fmt.Println("You must enter 2 arguments!")
	} else {
		mealTotal, _ := strconv.ParseFloat(args[0], 32)
		tipAmount, _ := strconv.ParseFloat(args[1], 32)
		fmt.Printf("\nYour meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
	}
}
func calculateTotal(mealTotal, tipAmount float32) float32 {
	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))
	return totalPrice
}
