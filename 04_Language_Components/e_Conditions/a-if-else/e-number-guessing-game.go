package main

import "fmt"

const LUCKYNUMBER int32 = 67

//Number Guessing Game
func main() {
	fmt.Println("Guess a number btwn 1 and 100:")

	var guessedNumber int32
	fmt.Scanf("%d", &guessedNumber)

	// Logic 1
	if guessedNumber == LUCKYNUMBER {
		println("You Guessed Correctly!!!")
	}

	// Logic 2
	if guessedNumber == LUCKYNUMBER {
		println("You Guessed Correctly!!!")
	} else {
		println("Please Try Again!!!")
	}

	// Logic 3
	if guessedNumber == LUCKYNUMBER {
		println("You Guessed Correctly!!!")
	} else if guessedNumber > LUCKYNUMBER { // 87 > 67
		println("Please reduce your guessing number")
	} else { // 56 < 67
		println("Please increase your guessing number")
	}
}
