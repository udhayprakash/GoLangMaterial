package main

import "fmt"

func linearSearch(dataList []int, searchVal int) bool {
	for _, item := range dataList {
		if item == searchVal {
			return true
		}
	}
	return false
}
func main() {
	items := []int{12, 34, 12, 78, 342, 54, 44}
	fmt.Println(12, linearSearch(items, 12))
	fmt.Println(67, linearSearch(items, 12))
}
