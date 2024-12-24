# GoDoc

    Golang documentation tool

## Installation

    go install golang.org/x/tools/cmd/godoc@latest

## Usage

    Go to the Project directory
        
        cd /path/to/your/project

    start godoc server 

        godoc -http :8080
        
    Open a browser and go to http://localhost:8080

    From there, project documentation can be navigates, similarly to https://pkg.go.dev 

    More importantly, when you click on a function, method, or type name, It will take you to the line where that function, method, variable, type, whatever, is defined.
