There are several performance check tools available in Golang that can be used to profile and optimize Go programs. Some of the most popular ones are:

pprof: This is a built-in Go tool that allows you to generate CPU and memory profiles and visualize them using the go tool pprof command. It is easy to use and provides detailed information about the performance of your program.

benchstat: This is a command-line tool that allows you to compare the performance of different versions of your program. It can be used to compare the results of multiple go test runs and identify changes in performance.

go-torch: This is a flame graph tool that allows you to visualize the performance of your program over time. It can be used to identify hot spots in your code and identify performance bottlenecks.

go-metrics: This is a library that allows you to collect and expose metrics from your Go program, such as CPU and memory usage, response times, and more. It can be used with various monitoring tools like Prometheus and Grafana to visualise and alert on the metrics.

goreplay: This is a tool that captures and replays live HTTP traffic, allowing you to reproduce and debug performance issues in production environments.

gops: This is a tool that allows you to diagnose and troubleshoot performance issues in running Go processes.

golangci-lint: This is a linter tool