package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Will print Number: 1234
	anInt := 1234
	fmt.Println("Number: " + strconv.Itoa(anInt))
}

// s = strconv.FormatInt(int64(i), 10) takes:  5.9133382s
// s = strconv.Itoa(int(i))            takes:  5.9763418s
// s = fmt.Sprint(i)                   takes: 13.5697761s
