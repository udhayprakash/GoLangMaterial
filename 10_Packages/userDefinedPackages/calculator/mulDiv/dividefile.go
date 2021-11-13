package muldiv

const Filename = "dividefile.go"

func Divide(n1, n2 float32) float32 {
	return n1 / n2
}

func DivideInts(n1, n2 int) float32 {
	return float32(n1) / float32(n2)
}
