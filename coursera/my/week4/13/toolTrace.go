package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}




func main() {

	runtime.GOMAXPROCS(4)
	http.HandleFunc("/", handle)

	fmt.Println("starting server at :8080")
	fmt.Println(http.ListenAndServe(" :8080", nil))
}



func handle(w http.ResponseWriter, r *http.Request) {

	result := ""

	for i := 0; i < 100; i ++  {
		currPost := &Post {ID: i, Text: "new post", Time: time.Now()
		}
		jsonRaw, _ := json.Marshal(currPost)
		result += string(jsonRaw)
	}
	time.Sleep(3 * time.Millisecond)
	w.Write([]byte(result))
}