package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}