package main

 import (
 	"fmt"
 	"regexp"
 )

 func main() {
 	str1 := "#$q4 revenue for 2017"

 	str2 := "revenue #$q2 for 2017"

 	str3 := "2017 revenue is #$q3"

 	// regular expression pattern
 	regE := regexp.MustCompile("\\#\\$q[1-4]")

 	fmt.Println(regE.FindAllString(str1, -1))
 	fmt.Println(regE.FindAllString(str2, -1))
 	fmt.Println(regE.FindAllString(str3, -1))

 	str4 := []string{str3, str1, str2}

 	for k, v := range str4 {
 		fmt.Println(k, regE.FindAllString(v, -1))
 	}

 }