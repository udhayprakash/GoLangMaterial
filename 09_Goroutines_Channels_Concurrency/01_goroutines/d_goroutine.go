package main

import (
	"fmt"
	"time"
)

func sleep_250(){
	fmt.Println("\nsleep_250 - start")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("sleep_250 - end")
}


func sleep_500(){
	fmt.Println("\nsleep_500 - start")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("sleep_500 - end")
}

func main(){
	fmt.Println("main -start")
	go sleep_250()
	go sleep_500()

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("main - end")
}
/*

BLOCKING EXECUTION
	main -start
	sleep_250 - start
	sleep_250 - end
	sleep_500 - start
	sleep_500 - end
	main - end

*/