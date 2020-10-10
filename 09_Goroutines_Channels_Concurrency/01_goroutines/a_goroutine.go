package main

import (
	"fmt"
	"time"
)


func printItem(i int) {
	fmt.Printf("Print Item: %2v\n", i + 1)
}

func main(){
	fmt.Println("Start Script")

	for j := 0; j < 10; j++{
		go printItem(j)
	}
	fmt.Println("End Script")

	// Wait, giving time for the go routines
	// to finish.
	time.Sleep(1000)

}
