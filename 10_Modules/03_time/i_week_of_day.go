package main

import (
	"fmt"
	"time"
)

func main() {

	//Sunday  Monday Tuesday Wednesday Thursday Friday Saturday
	// 0        1     2        3          4      5       6
	weekday := time.Now().Weekday()
	fmt.Println(weekday)      // "Friday"
	fmt.Println(int(weekday)) // 5

}
