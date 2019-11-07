package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
		)
	http.Handle("/data/", staticHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(" :8080", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte (`
	Hello <br/>
	<img src="/data/img/gopher.png" />
	`))
}


