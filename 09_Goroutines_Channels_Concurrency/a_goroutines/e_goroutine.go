package main

import (
	"fmt"
	"time"
)

func printNum(num int){
	fmt.Println("Giving number:", num)
}

func main(){
	fmt.Println("main - start")
	go printNum(1)
	go printNum(2)
	go printNum(3)

	time.Sleep(time.Second)
	fmt.Println("main - end")
}