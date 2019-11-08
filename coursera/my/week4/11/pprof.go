package main

import (
	"fmt"
	"net/http"
	_"net/http/pprof"
	"time"
)




func main() {

	http.HandleFunc("/", handle10)

	fmt.Println("starting server at :8080")
	fmt.Println(http.ListenAndServe(" :8080", nil))
}

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}



func handle10 (w http.ResponseWriter, req http.Request) {

	s := ""
	for i := 0; i < 1000; i++ {
		p := &Post{ID: i, Text: "new post"}
		s += fmt.Sprintf("%#v", p)
	}
	w.Write([]byte(s))
}