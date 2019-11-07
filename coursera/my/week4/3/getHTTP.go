package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(" :8080", nil)
}




func handler(writer http.ResponseWriter, request *http.Request) {

	myParam := request.URL.Query().
}
