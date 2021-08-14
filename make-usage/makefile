hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	go run main.go


compile:
	go build -o bin/main-windows.exe main.go
	echo "Compiling for every OS and Platform"
	set GOOS=linux
	set GOARCH=arm 
	go build -o bin/main-linux-arm main.go
	set GOOS=linux
	set GOARCH=arm64 
	go build -o bin/main-linux-arm64 main.go

all: hello build