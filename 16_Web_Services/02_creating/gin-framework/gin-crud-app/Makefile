# Variables
APP_NAME=gin-crud-app
DOCKER_IMAGE=gin-crud-app

# Build the application
build:
    go build -o bin/$(APP_NAME) ./cmd/main.go

# Run the application
run: build
    ./bin/$(APP_NAME)

# Run unit tests
test:
    go test -v ./tests/handlers/...

# Run integration tests
integration-test:
    go test -v ./tests/integration/...

# Run load tests with Vegeta (GET /groceries)
load-test-get:
    vegeta attack -duration=30s -rate=50 -targets=./tests/loadtests/targets-get.txt | vegeta report

# Run load tests with Vegeta (POST /groceries)
load-test-post:
    vegeta attack -duration=30s -rate=50 -targets=./tests/loadtests/targets-post.txt | vegeta report

# Build Docker image
docker-build:
    docker build -t $(DOCKER_IMAGE) .

# Run Docker container
docker-run: docker-build
    docker run -p 8080:8080 $(DOCKER_IMAGE)

# Clean up
clean:
    rm -rf bin/
    docker rmi $(DOCKER_IMAGE)