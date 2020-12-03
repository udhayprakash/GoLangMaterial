CPU profiling 
    go test -cpuprofile cpu.prof -bench .

Using pprof tool
    go tool pprof cpu.prof
        top5 -cum

Generating memory profiles
    go test -memprofile mem.prof -bench .

For generating both memory & CPU profiles
    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .