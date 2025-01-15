package main

import (
    "fmt"
    "time"
    vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
    // Define the target (request)
    targeter := vegeta.NewStaticTargeter(vegeta.Target{
        Method: "GET",
        URL:    "http://localhost:8080/groceries",
    })

    // Define the attack (load test)
    rate := vegeta.Rate{Freq: 50, Per: time.Second} // 50 requests per second
    duration := 30 * time.Second                    // Test duration
    attacker := vegeta.NewAttacker()

    // Run the attack
    var metrics vegeta.Metrics
    for res := range attacker.Attack(targeter, rate, duration, "Load Test") {
        metrics.Add(res)
    }
    metrics.Close()

    // Print the results
    fmt.Printf("Requests: %d\n", metrics.Requests)
    fmt.Printf("Success Rate: %.2f%%\n", metrics.Success*100)
    fmt.Printf("Latencies: %v\n", metrics.Latencies)
}