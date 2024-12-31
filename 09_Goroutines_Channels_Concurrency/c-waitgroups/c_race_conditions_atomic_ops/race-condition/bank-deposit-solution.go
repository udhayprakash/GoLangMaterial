package main

import (
	"fmt"
	"sync"
)

/*
Mutex
	- also known as mutual exclusion
	- prevents concurrent processes from accessing critical data
	  while a given process is performing another task.
*/

var (
	balance2 int
	wg2      sync.WaitGroup
	mu2      sync.Mutex
)

func Deposit2(amount int) {
	mu2.Lock()
	defer mu2.Unlock()

	balance2 += amount
	wg2.Done()
}

func main() {
	wg2.Add(3)
	go Deposit2(100)
	go Deposit2(200)
	go Deposit2(300)
	wg2.Wait()

	fmt.Println("Balance is:", balance2)
}

/*
OUTPUT:
------

	~go run bank-deposit-solution.go
	Balance is: 600

	~go run -race bank-deposit-solution.go
	Balance is: 600
*/
