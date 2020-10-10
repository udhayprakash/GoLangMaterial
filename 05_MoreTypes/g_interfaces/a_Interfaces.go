package main

import (
	"bytes"
	"fmt"
	"time"
)

func main(){
	var dozen int
	dozen = 12
	fmt.Println("dozen:", dozen)

	//dozen = 12.13  //  constant 12.13 truncated to integer

	var PI float64
	PI = 3
	fmt.Println("PI:", PI)

	PI = 3.1416
	fmt.Println("PI:", PI)

	//PI = "3.1416" // cannot use "3.1416" (type untyped string) as type float64 in assignment

	var any interface{}
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	any = true
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	any = 12.34
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	any = "hello"
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	any = map[string]int{"one": 1}
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	any = new(bytes.Buffer)
	fmt.Printf("any:%10v , dataType:%T\n", any, any)

	var x1 interface{} = []int{1, 2}
	fmt.Printf("x1 :%10v , dataType:%T\n", x1, x1)

	var x2 interface{} = time.Now()
	fmt.Printf("x2 :%10v , dataType:%T\n", x2, x2)
}
