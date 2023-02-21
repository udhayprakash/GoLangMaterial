package main

import (
	"container/list"
	"fmt"
)

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := list.New()

	for _, val := range vals {
		fmt.Printf("Pushing %v ... \n", val)
		fifo.PushBack(val)
	}
	fmt.Println("fifo =", fifo)

	for val := fifo.Front(); val != nil; val = val.Next() {
		fmt.Printf("Popped %v !\n", val.Value)
	}
}
