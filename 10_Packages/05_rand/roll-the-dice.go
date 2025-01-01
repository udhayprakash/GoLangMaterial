package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// prepare the dice
var dice = []int{1, 2, 3, 4, 5, 6}

var once sync.Once

func rollDice() int {

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(1 + rand.Intn(6-1))

	once.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})
	
	// fmt.Println(rand.IntN(6))
	// fmt.Println(rand.IntN(len(dice)))
	number := rand.Intn(len(dice))
	// fmt.Println(dice[number])

	return dice[number]
}

func main() {
	dice1 := rollDice()
	dice2 := rollDice()
	dice3 := rollDice()

	fmt.Println("Dice 1: ", dice1)
	fmt.Println("Dice 2: ", dice2)
	fmt.Println("Dice 3: ", dice3)
}
