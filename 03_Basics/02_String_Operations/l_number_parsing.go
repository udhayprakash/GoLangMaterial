package main

import (
	"fmt"
	"strconv"
)
func main(){
	fmt.Println(strconv.ParseFloat("1.234", 64))
	fmt.Println(strconv.ParseFloat("1.234 23", 64))

	value, err := strconv.ParseFloat("1.234", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	value, err = strconv.ParseFloat("1.2 34", 64)
	fmt.Println("Value = ", value, "\terror=", err)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
