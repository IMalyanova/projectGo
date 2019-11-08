package main

// шаблон в users.html

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}


func main() {

	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User {
		User {1, "Vasily", true},
		User {2, "<i>Ivan</i>", false},
		User {3, "Oleg", true},
	}

	http.HandleFunc("/", func(wr http.ResponseWriter, request *http.Request) {

		tmpl.Execute(w, struct {
			Users []User
		}{
			users,
		})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)

	//users.html находится сам шаблон
}
