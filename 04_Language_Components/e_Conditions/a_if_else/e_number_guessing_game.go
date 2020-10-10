package main

import "fmt"

const LUCKYNUMBER int32 = 67

func main() {
	//Number Guessing Game
	fmt.Println("Guess a number btwn 1 and 100:")

	var guessedNumber int32
	fmt.Scanf("%d", &guessedNumber)

	// Logic 1
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly")
	}

	// Logic 2
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly")
	} else {
		fmt.Println("Please Try again!!!")
	}

	// Logic 3
	if guessedNumber == LUCKYNUMBER {
		fmt.Println("You guessed correctly")
	} else if guessedNumber > LUCKYNUMBER { // 78 > 67
		fmt.Println("Please reduce your guessing number")
	} else {
		fmt.Println("Please increase your guessing number")
	}

}

/*
Assignment
		Get weekday name in runtime, and display the office timings of that day.
	Monday 		9 AM to 6 PM
	Tuesday		9 AM to 6 PM
	Wednesday	9 AM to 6 PM
	Thursday	9 AM to 6 PM
	Friday 		9 AM to 6 PM
	Saturday	9 AM to 1 PM
	Sunday 		HOLIDAY
*/
