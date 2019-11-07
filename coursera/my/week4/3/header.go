package main

import (
	"fmt"
	"net/http"
)



func main() {

	http.HandleFunc("/", handler3)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(" :8080", nil)
}




func handler3 (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("RequestID", "kds454vdvdfcv4854")

	fmt.Fprintln(w, "You browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}