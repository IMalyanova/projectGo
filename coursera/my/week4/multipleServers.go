package main

import (
	"fmt"
	"net/http"
)

func main() {

	go runServer(" :8081")
	runServer(" :8080")
}

func runServer(addr string) {

	mux := http.NewServeMux() // регистрируем новый обработчик запросов
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr:", addr, "URL:", r.URL.String())
		})

	server := http.Server {
		Addr:     addr,
		Handler:  mux,
	}

	fmt.Println("starting server at", addr)
	server.ListenAndServe()
}

