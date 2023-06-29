package main

import "fmt"

func main() {
	var userInput float32

	fmt.Printf("Please enter a floating point number: ")

	isSuccess, err := fmt.Scan(&userInput)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("err", err)


	if err != nil {
		fmt.Printf("An error occurred when trying to scan user input:\n\t%s\n", err)
	} else {
		fmt.Printf("The truncated integer value of your input is %d\n", int(userInput))
	}
}

/*

 $ go run  r_run_time_input.go 
Please enter a floating point number: 23
isSuccess 1
err <nil>
The truncated integer value of your input is 23

$ go run  r_run_time_input.go 
Please enter a floating point number: 23.4
isSuccess 1
err <nil>
The truncated integer value of your input is 23

$ go run  r_run_time_input.go 
Please enter a floating point number: age
isSuccess 0
err strconv.ParseFloat: parsing "": invalid syntax
An error occurred when trying to scan user input:
        strconv.ParseFloat: parsing "": invalid syntax

$ go run  r_run_time_input.go 
Please enter a floating point number: 23age
isSuccess 1
err <nil>
The truncated integer value of your input is 23

$ go run  r_run_time_input.go 
Please enter a floating point number: 23.4age
isSuccess 1
err <nil>
The truncated integer value of your input is 23

*/