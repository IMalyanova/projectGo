package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("starting server at :8082")
	http.ListenAndServe(":8082", nil)
}
