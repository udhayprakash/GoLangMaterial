package main

import (
	"fmt"
	"time"
)

var weekday string

func init() {
	fmt.Println("Init is called ...")
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Println("main is called ...")
	fmt.Printf("Today is %s", weekday)
}