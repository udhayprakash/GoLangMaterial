package main

import "fmt"

type Int int

func (i Int) Add(j Int) Int {
	return i + j
}

func main() {
	i := Int(5)
	j := Int(6)
	fmt.Println(i.Add(j))
	fmt.Println(i.Add(j) + 12)
}
