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

	for j := 10; j > 0; j-- {
		instance2 = GetInstance()
		fmt.Println("instance2=", instance2, &instance2)
	}

}

/*
~go run 02c_singleton.go
instance1= &{} 0xc000006028
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038
instance2= &{} 0xc000006038

*/
