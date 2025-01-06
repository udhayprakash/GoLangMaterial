Logging packages in Golang
---------------------------

    1) https://github.com/sirupsen/logrus 
        - used in many popular projects such as Docker
        -  Structured, pluggable logging for Go. (JSON and text formatting)
    
    2) https://github.com/uber-go/zap
        - from uber, used in high-performance applications
        - Blazing fast, structured, and leveled logging.
        - Optimized for performance, ideal for applications requiring low-latency logging.

    3) https://github.com/golang/glog 
        - from Google, implementation of their C++ glog library in Go

    4) https://github.com/go-kit/kit
        - popular in microservices and projects using go kit framework
        - focused on "structured logging" which is better for tools to consume
        - Lightweight and designed for structured logging in distributed systems.

    5) https://github.com/inconshreveable/log15
        - Used in projects like Geth (Go Ethereum).
        - Simple, structured logging with support for handlers and filters.

    6) https://github.com/op/go-logging 
        - smaller than the other here; used in opern source projects
        -  Supports different backends, levels, and formatting.

    7) https://github.com/rs/zerolog
        - Gaining popularity for its simplicity and performance.
        - Zero-allocation JSON logger, fast and efficient.
        - Ideal for performance-critical applications with minimal overhead.
        
    https://github.com/avelino/awesome-go#logging
