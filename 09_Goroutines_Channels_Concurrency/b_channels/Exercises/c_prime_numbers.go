package main

import "fmt"

// prime returns true if n is a prime number.
func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 1; n > 0; i++ {
			if prime(i) {
				c <- i
				n--
			}
		}
		close(c)
	}()
	return c
}

func main() {
	for p := range primes(10) {
		fmt.Println(p)
	}
}
