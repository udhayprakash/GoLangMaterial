# testing package


### testing.T

Purpose: Used for writing unit tests, benchmarks, and examples.

Key Methods:

    t.Error(args...) / t.Errorf(format, args...): Logs an error but continues test execution.

    t.Fatal(args...) / t.Fatalf(format, args...): Logs an error and stops test execution immediately.

    t.Log(args...) / t.Logf(format, args...): Logs informational messages during the test.

    t.Skip(args...) / t.Skipf(format, args...): Skips the test and marks it as skipped.

    t.Run(name string, f func(t *testing.T)): Runs a subtest.

    t.Parallel(): Marks the test as eligible for parallel execution.

    t.Helper(): Marks the calling function as a test helper (excludes it from error logs).


### testing.F

Purpose: Introduced in Go 1.18, it is used for writing fuzz tests (fuzz testing).

Key Methods:

    f.Fuzz(target func(T, ...)): Defines the fuzz target function to be tested with random inputs.

    f.Add(args...): Adds seed values to the corpus for fuzz testing.

    f.Skip(args...): Skips the fuzz test.


### testing.B:

Used for writing benchmarks.

Key methods: b.ResetTimer(), b.StartTimer(), b.StopTimer(), b.RunParallel()


### testing.M:

Used for setup and teardown of tests at the package level.

### testing.PB:

Used in parallel benchmarks with b.RunParallel().

### testing.TB:

An interface implemented by testing.T, testing.B, and testing.F. It provides common methods like Error, Fatal, Log, etc.

### testing.Coverage():

Returns the current code coverage as a fraction.

### testing.Short():

Used to check if the -short flag is set, allowing for skipping long-running tests.

### testing.AllocsPerRun():

Measures the number of allocations during a benchmark.


## Flags 

### General Test Execution Flags


-v (Verbose):

    Prints detailed output for each test, including logs and test names.

    Example: go test -v

-run:

    Runs only tests whose names match the provided regular expression.

    Example: go test -run TestAdd

-short:

    Skips long-running tests if they are guarded by testing.Short().

    Example: go test -short

-timeout:

    Sets a maximum duration for the test execution. If the tests exceed this duration, they are aborted.

    Example: go test -timeout 30s

-count:

    Runs each test the specified number of times. Useful for detecting flaky tests.

    Example: go test -count 5

-failfast:

    Stops test execution after the first failure.

    Example: go test -failfast

-list:

    Lists all tests matching the provided regex without running them.

    Example: go test -list TestAdd

-parallel:

    Sets the maximum number of tests to run in parallel.

    Example: go test -parallel 4

### Benchmark Flags


-bench:

    Runs benchmarks matching the provided regular expression.

    Example: go test -bench . (runs all benchmarks)

-benchtime:

    Sets the minimum time to run each benchmark.

    Example: go test -bench . -benchtime 5s

-benchmem:

    Includes memory allocation statistics in benchmark results.

    Example: go test -bench . -benchmem

-cpu:

    Specifies a list of GOMAXPROCS values to run benchmarks with.

    Example: go test -bench . -cpu 1,2,4


### Fuzz Testing Flags

-fuzz:

    Runs the fuzz test matching the provided regex.

    Example: go test -fuzz FuzzAdd

-fuzztime:

    Sets the duration for which to run fuzz tests.

    Example: go test -fuzz FuzzAdd -fuzztime 10s

-fuzzminimizetime:

    Sets the duration for which to minimize a failing input during fuzz testing.

    Example: go test -fuzz FuzzAdd -fuzzminimizetime 5s

### Code Coverage Flags

-cover:

    Enables code coverage analysis.

    Example: go test -cover

-coverprofile:

    Writes the coverage profile to a file.

    Example: go test -coverprofile=coverage.out

-covermode:

    Sets the coverage mode: set (default), count, or atomic.

    Example: go test -covermode=count

-coverpkg:

    Specifies a comma-separated list of packages to analyze for coverage.

    Example: go test -coverpkg=./pkg1,./pkg2

### Output and Logging Flags

-json:

    Outputs test results in JSON format.

    Example: go test -json

-o:

    Writes the test binary to the specified file.

    Example: go test -o test_binary

-trace:

    Writes an execution trace to the specified file.

    Example: go test -trace=trace.out

-vet:

    Configures the go vet tool to analyze the code during testing.

    Example: go test -vet=off (disables vetting)


### Miscellaneous Flags


-args:

    Passes additional arguments to the test binary.

    Example: go test -args -flag1=value1

-c:

    Compiles the test binary but does not run it.

    Example: go test -c

-i:

    Installs dependencies for the test but does not run it (deprecated in Go 1.16+).

-exec:

    Specifies an alternate command to run the test binary.

    Example: go test -exec "sudo"

-tags:

    Specifies build tags to control which files are included in the build.

    Example: go test -tags integration

-race:

    Enables race detection during testing.

    Example: go test -race

-msan:

    Enables memory sanitizer (requires C/C++ code and supported platforms).

-memprofile:

    Writes a memory profile to the specified file.

    Example: go test -memprofile=mem.out

-cpuprofile:

    Writes a CPU profile to the specified file.

    Example: go test -cpuprofile=cpu.out

-blockprofile:

    Writes a block profile to the specified file.

    Example: go test -blockprofile=block.out

-mutexprofile:

    Writes a mutex profile to the specified file.

    Example: go test -mutexprofile=mutex.out
