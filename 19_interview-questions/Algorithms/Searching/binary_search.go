package main

func binarySearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		if x == arr[m] {
			return m
		}
		if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}

func main() {
    items := []int{2, 3, 5, 7, 11, 13, 17}
	  
}
