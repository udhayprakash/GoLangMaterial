package main

import (
	"fmt"
)

func main() {
	num1 := 10
	fmt.Printf("Start - num1 = %d \t\t\t\taddr=%d\n", num1, &num1)
	{
		fmt.Printf("\tInside Start - num1 = %d \t\taddr=%d\n", num1, &num1)
		num1 := 20 // new local variable
		{
			fmt.Printf("\t\tInside Inside Start - num1 = %d addr=%d\n", num1, &num1)
			num1 := 30 // new local variable
			fmt.Printf("\t\tInside Inside End   - num1 = %d addr=%d\n", num1, &num1)
		}
		fmt.Printf("\tInside End   - num1 = %d \t\taddr=%d\n", num1, &num1)
	}
	fmt.Printf("End   - num1 = %d \t\t\t\taddr=%d\n", num1, &num1)
}

/*
OUTPUT:
	Start - num1 = 10                               addr=824633786536
			Inside Start - num1 = 10                addr=824633786536
					Inside Inside Start - num1 = 20 addr=824633786592
					Inside Inside End   - num1 = 30 addr=824633786600
			Inside End   - num1 = 20                addr=824633786592
	End   - num1 = 10                               addr=824633786536

*/
