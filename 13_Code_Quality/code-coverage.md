https://blog.seriesci.com/how-to-measure-code-coverage-in-go/

go tool cover -help

$ go test -cover .

    ok calc 0.001s coverage: 100.0% of statements

$ go test ./... -coverprofile cover.out

    ok calc 0.001s coverage: 100.0% of statements
    ok calc/format 0.001s coverage: 100.0% of statements

mocking

    https://blog.codecentric.de/en/2017/08/gomock-tutorial/


cheat-sheet: https://steveazz.xyz/blog/go-performance-tools-cheat-sheet/