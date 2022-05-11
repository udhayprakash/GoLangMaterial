you can install a tool called godoc 1 with this command:

go install golang.org/x/tools/cmd/godoc@latest
Then:

Change directory to your project
Run godoc -http :8080. Leave it running.
Open a browser and go to http://localhost:8080
From there, you can navigate the packages on your computer similarly to https://pkg.go.dev. More importantly, when you click on a function, method, or type name, It will take you to the line where that function, method, variable, type, whatever, is defined.
