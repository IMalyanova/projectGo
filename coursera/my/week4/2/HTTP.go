package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintln(w, "Hello!")
	w.Write([]byte("!!!"))
}


func main() {

	http.HandleFunc("/", handler)
	// HandleFunc вызывает базовый глобальный Multiplexor, который обрабатывает запросы

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
