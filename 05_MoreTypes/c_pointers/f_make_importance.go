package main

import "fmt"

func main() {
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})

	fmt.Println("ci=", ci)
	fmt.Println("cs=", cs)
	fmt.Println("cf=", cf)
}
