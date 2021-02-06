package main

import (
	"fmt"
	"time"
)

func sleep250(){
	fmt.Println("sleep250 - start")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("sleep250 - end")
}


func sleep500(){
	fmt.Println("sleep500 - start")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("sleep500 - end")
}

func main(){
	fmt.Println("main -start")
	go sleep250()
	go sleep500()

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("main - end")
}
/*

BLOCKING EXECUTION
	main -start
	sleep250 - start
	sleep250 - end
	sleep500 - start
	sleep500 - end
	main - end

*/