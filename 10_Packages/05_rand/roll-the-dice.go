package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once

// prepare the dice
var dice = []int{1, 2, 3, 4, 5, 6}

func rollDice() int {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(1 + rand.Intn(6-1))

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func main() {
	dice1 := rollDice()
	dice2 := rollDice()
	dice3 := rollDice()

	fmt.Println("Dice 1: ", dice1)
	fmt.Println("Dice 2: ", dice2)
	fmt.Println("Dice 3: ", dice3)
}
