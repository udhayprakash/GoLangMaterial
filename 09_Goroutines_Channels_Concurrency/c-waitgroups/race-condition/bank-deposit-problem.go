package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	wg      sync.WaitGroup
)

func Deposit(amount int) {
	balance += amount
	wg.Done()
}

func main() {
	wg.Add(3)
	go Deposit(100)
	go Deposit(200)
	go Deposit(300)
	wg.Wait()

	fmt.Println("Balance is:", balance)
}

// go run -race bank-deposit-example.go
