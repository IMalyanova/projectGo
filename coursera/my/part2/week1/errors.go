package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}



func handler(w http.ResponseWriter, r *http.Request) {

	err := getRemoteResource()

	if err != nil {
		fmt.Println("error happend: %+v\n", err)
		http.Error(w, "internal error", 500)
		return
	}

	w.Write([]byte("all is OK"))
}

var client = http.Client{
	Timeout: time.Duration(time.Millisecond),
}

func getRemoteResource() error {

	url := "http://127.0.1:9999/pages?id=123"
	_, err := client.Get(url)

	if err != nil {
		return fmt.Errorf("getRemoteResource: %s at %s", err, url)
	}
	return nil
}
