// Parse below string, without using any package,
// and display keys below 5
//  A%1,B%2,C%13,D:4,E:5,F:6
package main

import "fmt"

func main() {
	data := "A%1,B%2,C%13,d:4,e:5,F:2"
	temp := ""
	var result = map[string]string{}
	for _, char := range data {
		println("char:", char, string(char))
		if string(char) != "%" && string(char) != ":" && string(char) != "," {
			if temp == "" {
				temp = string(char)
			} else {
				if char <= 53 {
					// fmt.Println(temp, string(char))
					result[temp] = string(char)
				}
				temp = ""
			}
		}
	}
	fmt.Println("result=", result) // map[A:1 B:2 C:1]
}
