package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, err := template.New("todos").Parse("\nYou have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	// reusing existing templates
	tdNew := Todo{"Go", "Contribute to any Go project"}
	err = t.Execute(os.Stdout, tdNew)
}

/*
Parsing Templates:
------------------
	New 	— allocates new, undefined template,
	Parse 	— parses given template string and return parsed template,
	Execute — applies parsed template to the data structure and writes result to the given writer.

*/
