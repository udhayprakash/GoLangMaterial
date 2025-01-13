package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
	Quote      string `json:"quote"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulate fetching data from JSON file
		data, err := ioutil.ReadFile("data/persons.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var persons []Person
		err = json.Unmarshal(data, &persons)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		person := persons[rand.Intn(len(persons))] // Select a random person

		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/profile.html"))
		if err := tmpl.ExecuteTemplate(w, "profile.html", person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
