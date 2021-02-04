package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	s := new(runtime.MemStats)
	runtime.ReadMemStats(s)

	fmt.Println("MEMORY STATS")
	fmt.Println("Alloc       : ", s.Alloc)
	fmt.Println("Total Alloc : ", s.TotalAlloc)
	fmt.Println("Sys         : ", s.Sys)
	fmt.Println("Lookups     : ", s.Lookups)

	fmt.Println("\nGarbage Collection data")
	fmt.Println("HeapAlloc               : ", s.HeapAlloc)
	fmt.Println("Next Garbage Collection : ", s.NextGC)
	fmt.Println("Auto garbage collection will happen when HeapAlloc >= NextGC")
	fmt.Println("Last Garbage Collection   : ", s.LastGC)
	fmt.Println("Total number of GC pause  : ", s.PauseTotalNs)

	// [256]uint64 array
	// we will take the first element for this example...you might want to loop the array
	fmt.Println("\nMost recent pause         : ", s.PauseNs[0])
	fmt.Println("Recent pause end times    : ", s.PauseEnd[0])

	fmt.Println("Number of Garbage Collections  : ", s.NumGC)
	fmt.Println("Is Garbage Collection enabled? : ", s.EnableGC)

	fmt.Println("Is Garbage Collection debug enabled? : ", s.DebugGC)

	// ---------------------------

	gcs := new(debug.GCStats)
	debug.ReadGCStats(gcs)

	fmt.Println("Last Garbage Collection         : ", gcs.LastGC)

	fmt.Println("Number of Garbage Collection    : ", gcs.NumGC)
	fmt.Println("Total pause for all collections : ", gcs.PauseTotal)

	// []time.Duration array
	fmt.Println("Most recent pause history       : ", gcs.Pause)

	// []time.Time array
	fmt.Println("Most recent pause end times history : ", gcs.PauseEnd)

	// []time.Duration array
	// Pause Quantiles == values taken at regular interval
	fmt.Println("Most recent pause quantiles history : ", gcs.PauseQuantiles)
}
