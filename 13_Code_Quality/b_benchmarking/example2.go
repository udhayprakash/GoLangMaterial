package benchmarking 

func IntAdd(words [][]byte) map[string]int {
    var m = make(map[string]int)
    for _, w := range words {
        m[string(w)] = m[string(w)] + 1
       
    }
    return m
}

func IntIncrement(words [][]byte) map[string]int {
    var m = make(map[string]int)
    for _, w := range words {
        m[string(w)]++
    }
    return m
}

// IntAdd is obviously slower than IntIncrement.