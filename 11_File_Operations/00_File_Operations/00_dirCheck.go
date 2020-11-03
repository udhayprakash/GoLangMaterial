package main

import (
	"fmt"
	"os"
)

func main(){
	// To create a new directory
	output := os.Mkdir("MyDIR", 755)
	fmt.Println("output=", output)

	dirInfo, err:= os.Stat("MyDIR")
	if err != nil{
		fmt.Println("err:", err)
	}
	fmt.Println("dirInfo:", dirInfo)

	cwd, err:= os.Getwd()
	if err != nil{
		fmt.Println("err:", err)
	}
	fmt.Println("current working dir:", cwd)

	// changing the current working directory
	os.Chdir("MyDIR")
	cwd, err = os.Getwd()
	if err != nil{
		fmt.Println("err:", err)
	}
	fmt.Println("current working dir:", cwd)
}
