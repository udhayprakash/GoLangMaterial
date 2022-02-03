package mergesort

func merge(a []int, b []int) []int {
	var r = make([]int, len(a)+len(b))
	var i = 0

	for i < len(a) && i < len(b) {
		if a[i] <= b[i] {
			r[i+1] = a[i]
			i++
		} else {
			r[i+1] = b[i]
			i++
		}
	}

	for i < len(a) {
		r[i+1] = a[i]
		i++
	}
	for i < len(b) {
		r[i+1] = b[i]
		i++
	}
	return r
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	var middle = len(items)
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)
}
