package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Calculation struct {
	Num1     float64
	Num2     float64
	Operator string
	Result   float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/calculator.html"))
		if r.Method == http.MethodPost {
			num1Str, num2Str := r.FormValue("num1"), r.FormValue("num2")
			operator := r.FormValue("operator")
			num1, _ := strconv.ParseFloat(num1Str, 64)
			num2, _ := strconv.ParseFloat(num2Str, 64)
			var result float64
			switch operator {
			case "+":
				result = calculateSum(num1, num2)
			case "-":
				result = calculateSubtraction(num1, num2)
			case "*":
				result = calculateMultiplication(num1, num2)
			case "/":
				result = calculateDivision(num1, num2)
			default:
				result = -1 // Invalid operation
			}
			calculation := Calculation{
				Num1:     num1,
				Num2:     num2,
				Operator: operator,
				Result:   result,
			}
			tmpl = template.Must(template.ParseFiles("templates/base.html", "templates/result.html"))
			if err := tmpl.ExecuteTemplate(w, "result.html", calculation); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		if err := tmpl.ExecuteTemplate(w, "calculator.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func calculateSum(num1, num2 float64) float64 {
	return num1 + num2
}

func calculateSubtraction(num1, num2 float64) float64 {
	return num1 - num2
}

func calculateMultiplication(num1, num2 float64) float64 {
	return num1 * num2
}

func calculateDivision(num1, num2 float64) float64 {
	if num2 == 0 {
		return -1 // Avoid division by zero
	}
	return num1 / num2
}
