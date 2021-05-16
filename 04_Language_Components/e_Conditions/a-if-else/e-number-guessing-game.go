package main

import "fmt"

const LUCKYNUMBER = 67

// Number Guessing Game
func main() {
	fmt.Println("Guess a number btwn 1 and 100:")

	var guessedNumber int
	fmt.Scanf("%d", &guessedNumber)
	// fmt.Println("guessedNumber=", guessedNumber)

	// Logic 1
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly!!!")
	}

	// Logic 2
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly!!!")
	} else {
		println("Please Try Again!!!")
	}

	// Logic 3
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly!!!")
	} else if guessedNumber > LUCKYNUMBER { // 89 > 67
		fmt.Println("Please reduce your guessing number")
	} else { // 56 < 67
		fmt.Println("Please increase your guessing number")
	}
}
