package main

import "fmt"

func main() {
	fruit := "apple"
	fmt.Println("fruit:", fruit)

	// NOTE 1: keywords should not be used as identifiers
	// break := "one" //syntax error: unexpected := at end of statement

	// NOT RECOMMENDED - predeclared variables should not be used as keywords
	true := "One"
	fmt.Println("true:", true)

	break_one := "one"
	fmt.Println("break_one =", break_one)

	// NOTE 2: CamelCase is recommended for identifiers

	breakOne := "one"
	fmt.Println("breakOne=", breakOne)

	costOfMangos := 34
	fmt.Println("costOfMangos=", costOfMangos)

	NoOfProcessesRunning := 3
	fmt.Println("NoOfProcessesRunning=", NoOfProcessesRunning)

}
