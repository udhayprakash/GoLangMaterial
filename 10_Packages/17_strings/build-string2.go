package main

import (
	"fmt"
	"strings"
)

func main() {

	builder := strings.Builder{}

	data1 := []byte{72, 101, 108, 108, 111}
	data2 := []byte{32}
	data3 := []byte{116, 104, 101, 114, 101, 33}

	builder.Write(data1)
	builder.Write(data2)
	builder.Write(data3)

	fmt.Println(builder.String())
}
