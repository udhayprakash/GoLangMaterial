package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	m := &sync.Mutex{}

	var answer string = ""
	fmt.Print("do you want to continue?: ")
	
	go func() {
		time.Sleep(5 * time.Second)

		m.Lock()
		isEmpty := answer == ""
		m.Unlock()
		
		if isEmpty {
			fmt.Println("timeout")
			os.Exit(0)
		}
	}()

	fmt.Scanln(&answer)
	fmt.Println("answer is " + answer)
}
