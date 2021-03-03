CPU profiling go test -cpuprofile cpu.prof -bench .

Using pprof tool go tool pprof cpu.prof top5 -cum

Generating memory profiles go test -memprofile mem.prof -bench .

For generating both memory & CPU profiles go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

profilers - Go has several built in profiles:
- Goroutine: stack traces of all current Goroutines - CPU: stack traces of CPU returned by runtime - Heap: a sampling of
memory allocations of live objects - Allocation: a sampling of all past memory allocations - Thread: stack traces that
led to the creation of new OS threads - Block: stack traces that led to blocking on synchronization primitives - Mutex:
stack traces of holders of contended mutexes - using "go test"
- go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
