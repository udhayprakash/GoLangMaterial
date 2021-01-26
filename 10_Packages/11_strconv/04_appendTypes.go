package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b)) // bool:true

	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32)) // float32:3.1415927E+00

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64)) // float64:3.1415926535E+00

	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) // int (base 10):-42

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) // int (base 16):-2a

	bq := []byte("quote:")
	bq = strconv.AppendQuote(bq, `"Fran & Freddie's Diner"`)
	fmt.Println(string(bq)) // quote:"\"Fran & Freddie's Diner\""

	br := []byte("rune:")
	br = strconv.AppendQuoteRune(br, '☺')
	fmt.Println(string(br)) // rune:'☺'

	ba := []byte("rune (ascii):")
	ba = strconv.AppendQuoteRuneToASCII(ba, '☺')
	fmt.Println(string(ba)) // rune (ascii):'\u263a'

}
