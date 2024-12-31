package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum of numbers:", sum(numbers))
}

func sum(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

/*
start debugger 

	dlv debug main.go

To show current executing line 
	list

set breakpoint

	break sum
	break b_with_function_call.go:10
	break b_with_function_call.go:13
	break b_with_function_call.go:15

List all breakpoints

	breakpoints

continue to next steps 

	continue


clear breakpoints
	clear 1		// to clear assigned breakpoints, based on order of assignment (not line number), from 1
	clearall    // to clear all breakpoints

Step through the loop:

	step

Exit the debugger

	quit


-----------------------------------------------------------------
Command			Purpose
-----------------------------------------------------------------
set			Modify the value of a simple type variable.
call		Execute a method, function, or complex modification 
			(e.g., modify a slice element).
-----------------------------------------------------------------

GOTCHA
=======
	(dlv) set nums  []int{1, 2, 3, 4, 5, 6}
	Command failed: 1:8: expected operand, found ']'

	// NOTE: we can set complex type variables with dlv

	workaround
	-----------
		call nums[0] = 1
		call nums[1] = 2
		call nums[2] = 3
		call nums[3] = 4
		call nums[4] = 5


	(dlv) call nums[5] = 6
	> main.sum() ./b_with_function_call.go:11 (PC: 0x4af07f)
	Command failed: index out of bounds

	// NOTE: Also, we cant change the size, dynamically, even if it is slice
*/


// dlv debug filename.go
