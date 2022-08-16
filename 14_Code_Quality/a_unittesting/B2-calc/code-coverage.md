Since version 1.2 code coverage is built into Go (Golang). Use go test -cover . to get basic coverage statistics for a
single package.

    $ go test -cover .
    ok      calc        0.001s  coverage: 100.0% of statements

Unfortunately it does not show the total coverage for all your packages within a single project.

To get the overall code coverage for multiple packages,

    $ go test ./... -coverprofile cover.out
    ok      calc        0.001s  coverage: 100.0% of statements
    ok      calc/format 0.001s  coverage: 100.0% of statements

Afterwards use go tool cover with the -func flag and point it to the previously generated file. This will show you the
code coverage for every single package withing your project and down at the bottom the total coverage.

    $ go tool cover -func cover.out
    calc/calc.go:4:              Add             100.0%
    calc/calc.go:9:              Subtract        100.0%
    calc/calc.go:14:             Multiply        100.0%
    calc/format/format.go:8:     Print           100.0%
    total:                       (statements)    100.0%