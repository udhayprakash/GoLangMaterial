# Property Based Testing


## comparision

Testing (testing.T): 
    
    Best for verifying specific functionality.
    Example: Verify that Add(2, 3) returns 5.


Quick Testing (testing/quick):
    
    Best for verifying general properties or invariants.
    Example: Verify that Add(a, b) == Add(b, a) for random a and b.


Benchmarking (testing.B):
    
    Best for measuring performance.
    Example: Measure how long Add(2, 3) takes to execute.


Fuzz Testing (testing.F):
    
    Best for finding crashes or edge cases.
    Example: Test Add(a, b) with random a and b to find edge cases.




| **Aspect**              | **Testing (`testing.T`)**                  | **Quick Testing (`testing/quick`)**     | **Benchmarking (`testing.B`)**          | **Fuzz Testing (`testing.F`)**          |
|-------------------------|--------------------------------------------|-----------------------------------------|-----------------------------------------|-----------------------------------------|
| **Purpose**             | Verify specific input-output pairs         | Verify properties for random inputs     | Measure performance and resource usage  | Find crashes or unexpected behavior     |
| **Input Generation**    | Manual (fixed inputs)                      | Random (customizable via generators)    | Fixed (no input generation)             | Random (based on seed corpus)           |
| **Execution**           | Runs once per test case                    | Runs fixed iterations (default: 100)    | Runs repeatedly (`b.N` times)           | Runs until crash or timeout             |
| **Stopping Condition**  | Stops on failure or completion             | Stops after `MaxCount` or property fail | Stops after `b.N` iterations            | Stops on crash or after `-fuzztime`     |
| **Use Cases**           | Unit tests, integration tests              | Verify invariants, logical properties   | Measure execution time, memory usage    | Find crashes, panics, edge cases        |
| **Customization**       | Limited (fixed inputs)                     | High (custom generators)                | None                                    | Medium (seed corpus, fuzz functions)    |
| **Output**              | Pass/fail for specific test cases          | Pass/fail based on property             | Performance metrics (ns/op, B/op, etc.) | Crash report or failure                 |
| **Example**             | ```go func TestAdd(t *testing.T) { ... }```| ```go quick.Check(func(a, b int) bool { return Add(a, b) == Add(b, a) }, nil)``` | ```go func BenchmarkAdd(b *testing.B) { for i := 0; i < b.N; i++ { Add(2, 3) } }``` | ```go func FuzzAdd(f *testing.F) { f.Fuzz(func(t *testing.T, a int, b int) { ... }) }``` |
| **Key Methods**         | `t.Error`, `t.Fatal`, `t.Log`, `t.Run`     | `quick.Check`, `quick.CheckEqual`       | `b.ResetTimer`, `b.StartTimer`, `b.StopTimer` | `f.Fuzz`, `f.Add`, `f.Skip`             |
| **When to Use**         | Test specific functionality                | Test mathematical or logical properties | Optimize performance                    | Find edge cases or crashes              |