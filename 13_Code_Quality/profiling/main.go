package main
 
// func Fib2(n int) uint64 {
//     if n == 0 {
//         return 0
//     } else if n == 1 {
//         return 1
//     } else {
//         return Fib2(n-1) + Fib2(n-2)
//     }
// }
 
var f [100010]uint64
 
func Fib2(n int) uint64 {
    f[0] = 0
    f[1] = 1
 
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
 
    return f[n]
}


// // For memory profiling
// func Fib(n int) uint64 {
 
//     f := make([]uint64, 10000)       // make huge allocation
 
//     f[0] = 0
//     f[1] = 1
 
//     for i := 2; i <= n; i++ {
//         f[i] = f[i-1] + f[i-2]
//     }
 
//     return f[n]
// }

func main() {
    // fmt.Println(Fib2(30)) // 832040
}