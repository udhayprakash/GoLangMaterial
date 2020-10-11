package main

import "fmt"

func main(){

	fmt.Println("\n\nFull Loop")
	for i:= 0; i<10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("\n\nImportance of break")
	for i:= 0; i<10; i++ {
		if i==5 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println("\n\nImportance of continue")
	for i:= 0; i<10; i++ {
		if i==5 {
			continue
		}
		fmt.Print(i, " ")
	}
}
/*
NOTE:
A continue statement begins the next iteration of the
innermost for loop at its post statement.
A break statement terminates execution of the innermost
for, switch, or select statement.
*/
