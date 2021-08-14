package main 

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct{
	side float32
}
func (sq *Square) Area() float32{
	return sq.side * sq.side
}

func main(){
	square := new(Square)
	square.side = 10
	areaIntf := square 
	fmt.Printf("Square area: %f\n", areaIntf.Area())
}