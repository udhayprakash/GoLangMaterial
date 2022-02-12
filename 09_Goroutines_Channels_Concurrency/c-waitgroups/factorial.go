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

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go factorial(i, &wg)
	}

	// wg.Add(1)
	// go factorial(4, &wg)

	wg.Wait()
	fmt.Println("factorials = ", factorials)
}

/*
OUTPUT:
-------
	factorial(5) = 120
	factorial(2) = 2
	factorial(1) = 1
	factorial(4) = 24
	factorial(3) = 6
	factorials =  map[1:1 2:2 3:6 4:24 5:120]

*/
