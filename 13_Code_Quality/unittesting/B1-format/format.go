package format

import "strconv"

func Print(value int, decimals int) string {
	return strconv.FormatFloat(float64(value), 'f', decimals, 64)
}
