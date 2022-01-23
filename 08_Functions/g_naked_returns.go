package main

import "fmt"

/*
A return statement without arguments returns the named return
values. This is known as a "naked" return.

Naked return statements should be used only in short functions,
as with the example shown here. They can harm readability in
longer functions.

Also, called as named return

*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// Just return keyword, without any resultant values
	return
}

func main() {
	fmt.Println(split(17))
}
