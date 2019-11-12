package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", List)
	r.HandleFunc("/users", List).Host("localhost")
	r.HandleFunc("/users", Create).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", Get)
	r.HandleFunc("/users/{login}", Update).Methods("POST").Headers("X-Auth", "test")

	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(" :8080", r))
}



func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to update %s\n", vars["login"])
}



func Create(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "you try to create new user\n")
}



func List (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You see user list\n")
}



func Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to see user %s\n", vars["id"])
}