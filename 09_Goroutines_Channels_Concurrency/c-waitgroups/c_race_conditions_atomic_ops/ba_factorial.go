// factorial(9) = 9 * 8 * 7 * 6 * ... 1
package main

import (
	"fmt"
	"sync"
)

var factorials = make(map[int]int)

func factorial(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}

	var wg2 sync.WaitGroup

	wg2.Add(1)
	go DisplayResult(num, res, &wg2)
	wg2.Wait()

}

func DisplayResult(n, r int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	fmt.Printf("factorial(%d) = %d\n", n, r)
	factorials[n] = r
}

func main() {
	wg := sync.WaitGroup{}

	// fmt.Println("factorial(10) = ", factorial(10))

	for i := 1; i < 5; i++ {
		wg.Add(1)
		// go fmt.Printf("\tfactorial(%d, &wg) = %d\n", i, factorial(i, &wg))
		go factorial(i, &wg)
	}

	wg.Wait()
	fmt.Println("factorials = ", factorials)

}

/*
OUTPUT IS INCONSISTENT, because of race condition

$ go run factorial.go
factorial(1) = 1
factorial(4) = 24
factorial(2) = 2
factorials =  map[1:1]

$ go run factorial.go
factorial(4) = 24
factorial(1) = 1
factorial(2) = 2
factorials =  map[]

$ go run factorial.go
factorial(4) = 24
factorial(1) = 1
factorial(2) = 2
factorials =  map[1:1 2:2 4:24]


*/
