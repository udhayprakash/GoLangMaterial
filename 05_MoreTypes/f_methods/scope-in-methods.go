package main 


type Cat struct {
	IsSleeping bool
}

func (c Cat) Sleep () {
	c.IsSleeping = true
}

func main(){
	billy := Cat{}
	billy.Sleep()
	println(billy.IsSleeping) // false
}