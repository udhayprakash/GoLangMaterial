package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Game settings
const (
	maxAttempts = 5
	timeout     = 10 * time.Second
)

func main() {
	// Generate a random number between 1 and 100
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	target := r1.Intn(100) + 1
	fmt.Println("Welcome to the Guessing Number Game!")
	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Start the game
	attempts := 0
	for {
		select {
		case <-ctx.Done():
			// Handle timeout or cancellation
			fmt.Println("\nGame over! Time's up or you canceled the game.")
			fmt.Printf("The correct number was: %d\n", target)
			return
		default:
			// Prompt the player for a guess
			fmt.Printf("\nAttempt %d/%d: Enter your guess: ", attempts+1, maxAttempts)
			var guess int
			_, err := fmt.Scanf("%d", &guess)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}

			// Check the guess
			attempts++
			if guess < target {
				fmt.Println("Too low!")
			} else if guess > target {
				fmt.Println("Too high!")
			} else {
				fmt.Println("Congratulations! You guessed the number!")
				return
			}

			// Check if the player has used all attempts
			if attempts >= maxAttempts {
				fmt.Println("\nGame over! You've used all your attempts.")
				fmt.Printf("The correct number was: %d\n", target)
				return
			}
		}
	}
}
