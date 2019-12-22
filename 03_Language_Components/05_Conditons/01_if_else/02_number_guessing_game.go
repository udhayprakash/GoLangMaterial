package main

import "fmt"

const LUCKY_NUMBER = 69

func main() {
	//Number Guessing Game
	fmt.Println("Guess a number btwn 1 and 100:")

	var guessed_number float64
	fmt.Scanf("%f", &guessed_number)

	if guessed_number == LUCKY_NUMBER {
		fmt.Println("YOu Guessed correctly!")
	} else if guessed_number > LUCKY_NUMBER {
		fmt.Println("Reduce your guess!")
	} else {
		fmt.Println("Increase your guess!")
	}

}
