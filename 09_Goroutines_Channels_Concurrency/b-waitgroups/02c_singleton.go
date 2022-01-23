package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	instance1 := GetInstance()
	fmt.Println("instance1=", instance1, &instance1)

	instance2 := GetInstance()
	fmt.Println("instance2=", instance2, &instance2)
}
