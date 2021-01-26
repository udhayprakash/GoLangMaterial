package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var str string = "123456789"

	sixtyFour, err := strconv.ParseInt(str, 10, 64) // use base 10 for sanity

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The type of str is %v \n", reflect.TypeOf(str))

	fmt.Printf("The type of sixtyFour is %v \n", reflect.TypeOf(sixtyFour))

}
