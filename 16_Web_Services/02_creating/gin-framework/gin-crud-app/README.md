# Grocery Store REST API

## Installation
    
    github.com/gin-gonic/gin
    go get github.com/golang-jwt/jwt/v5

## Project Setup

    $ go mod init gin-crud-app
    $ go mod tidy

### Project Folder Structure

    gin-crud-app/
    ├── cmd/
    │   └── main.go                     # Entry point of the application
    ├── internal/
    │   ├── config/
    │   │   └── config.go               # Configuration setup (e.g., environment variables)
    │   ├── auth/
    │   │   ├── auth.go                 # Authentication logic (JWT, login)
    │   │   └── middleware.go           # Authentication and authorization middleware
    │   ├── handlers/
    │   │   └── grocery.go              # HTTP handlers for grocery routes
    │   ├── models/
    │   │   ├── user.go                 # User model and related logic
    │   │   └── grocery.go              # Grocery model and related logic
    │   ├── services/
    │   │   └── grocery.go              # Business logic for grocery operations
    │   └── utils/
    │       └── utils.go                # Utility functions (e.g., error handling)
    ├── pkg/
    │   └── database/                   # Database connection and setup (if applicable)
    ├── tests/
    │   ├── handlers/
    │   │   └── grocery_test.go         # Unit tests for grocery handlers
    │   ├── loadtests/
    │   │   ├── loadtest.js             # k6 load test script
    │   │   ├── targets-get.txt         # Vegeta targets for GET /groceries
    │   │   ├── targets-post.txt        # Vegeta targets for POST /groceries
    │   │   └── loadtest.go             # Programmatic load test using Vegeta
    │   └── integration/
    │       └── integration_test.go     # Integration tests for the application
    ├── Dockerfile                      # Dockerfile for containerization
    ├── Makefile                        # Makefile for automation
    ├── go.mod                          # Go module file
    ├── go.sum                          # Go checksum file
    └── README.md                       # Project documentation


## Run application

    $ go run cmd/main.go

## Access the API

    Access the API at http://localhost:8080.


## Testing

## Unit testing 

    $ go test -v

## Load testing

    $ GET http://localhost:8080/groceries



## Make file Usage

1) Build the application:

    make build

2) Run the application:

    make run

3) Run unit tests:

    make test

4) Run integration tests:

    make integration-test

5) Run load tests for GET /groceries:

    make load-test-get

6) Run load tests for POST /groceries:

    make load-test-post

7) Build and run the Docker container:

    make docker-run

8) Clean up:

    make clean

