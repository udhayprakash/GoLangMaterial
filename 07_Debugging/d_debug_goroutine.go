package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Goroutine:", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

/*


dlv debug d_debug_goroutine.go
break d_debug_goroutine.go:11
continue
goroutines // to list all active goroutines
goroutine <id>  // switch to a specific goroutine
step  // step through the goroutine
quit  // Exit the debugger

n line by line
step block by block
continue -- breakpoint wise



goroutines
* Goroutine 1 - User: ./d_debug_goroutine.go:11 main.main (0x4af37f) (thread 16183)
  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:425 runtime.gopark (0x46dc7c) [force gc (idle)]
  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:425 runtime.gopark (0x46dc7c) [finalizer wait]
  Goroutine 4 - User: ./d_debug_goroutine.go:13 main.main.gowrap1 (0x4af420)
  Goroutine 17 - User: /usr/local/go/src/runtime/proc.go:425 runtime.gopark (0x46dc7c) [GC sweep wait]
  Goroutine 18 - User: /usr/local/go/src/runtime/proc.go:425 runtime.gopark (0x46dc7c) [GC scavenge wait]
[6 goroutines]


Builtin goroutines
 runtime.gopark
	User: /usr/local/go/src/runtime/proc.go
	
	[force gc (idle)]
	// This goroutine is waiting for a garbage collection (GC) cycle. It is part of Go's automatic memory management.

	[finalizer wait]
	// This goroutine is used to handle finalizers for objects that require cleanup before garbage collection.

	[GC sweep wait]
	// This goroutine waits for a GC sweep operation.

	[GC scavenge wait]
	// This goroutine waits for a GC scavenge operation.




*/
