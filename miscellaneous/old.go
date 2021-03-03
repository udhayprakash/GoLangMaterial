package main

// func main() {
// 	// i, j := 42, 2701
// 	// p, q := &i, &j
// 	// *p, *q = 21, 42
// 	// fmt.Println(*p, *q, i, j)
// 	i, j, k, _ := 1, 2, 3, 4
// 	i, j, k = j, i, k
// 	_, k, _ = i, j, j
// 	fmt.Println(i, j, k)
// }

// func main(){
// 	i := 0
// 	n := 1
// 	for i =0; i <5; i++{
// 		n += i *2
// 		i = n-1
// 	}
// 	fmt.Printf("%d %d \n ", i, n)
// }

// func main(){
// 	f := func (i int, s []int){
// 		s = append(s, i)
// 		fmt.Printf("in %v \n", s)
// 	}

// 	src := make([]int, 4, 5)
// 	src[0] = 1
// 	src[1] = 1
// 	f(2, src[:3])
// 	f(3, src[:5])
// 	fmt.Printf("%v\n", src)
// }

// func main() {
// 	buf := []byte("String")
// 	for i := 0; i < len(buf); i++ {
// 		go func(i int) {
// 			fmt.Printf("%c", buf[i])
// 		}(i)
// 	}
// 	time.Sleep(1000)
// }

// func main() {
// 	var s string
// 	l := "SrcNoR"
// 	// for 1
// 	for i := 0; i < utf8.RuneCountInString(l); i++ {
// 		s = s + string(l[i])
// 	}
// 	fmt.Println("for 1", s)

// }

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 1
// 	}()

// 	fmt.Printf(" %d \n", <-c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6, 7}
// 	n := s[3:5]
// 	_ = append(n, 10)
// 	_ = append(n, 20)
// 	_ = append(n, 30)
// 	n[0] = 40
// 	fmt.Printf("%v", s)

// }

// func exec1() {
// 	defer fmt.Println("exec 1 begin")
// 	exec2()
// 	defer fmt.Println("exec 1 finish")
// }

// func exec2() {
// 	defer fmt.Println("exec 2 begin")
// 	exec3()
// 	defer fmt.Println("exec 2 finish")
// }

// func exec3() {
// 	defer fmt.Println("exec 3 begin")
// 	fmt.Println("panic(100)")
// 	defer fmt.Println("exec 3 finish")
// }

// func main() {
// 	exec1()
// }

// func main() {
// 	// var m map[int]int
// 	n := make([]int, 5)

// 	var s []string

// 	s = append(s, "a")
// 	// s[1] = "b"
// 	// m[0] = 0
// 	n[0] = 6
// 	s[0] = "c"

// 	s = append(s, string(n[3]))
// }

func main() {

	var prod map[string]int
	prod["a"] = 1
	prod["p"] = 0

	//if v := prod["p"]
}
