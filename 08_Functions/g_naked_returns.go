package main

import "fmt"

/*
A return statement without arguments returns the named return
values. This is known as a "naked" return.

Naked return statements should be used only in short functions,
as with the example shown here. They can harm readability in
longer functions.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
