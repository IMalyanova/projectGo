package main

import (
	"bytes"
	"coursera/template_adv/item"
	"coursera/template_adv/template"
	"fmt"
	"net/http"
)

var ExampleItems = []*item.Item {
	&item.Item {1, "rvasily", "Mail.ru Group"},
	&item.Item {2, "username", "freelancer"},
}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r http.Request) {
		buffer := new(bytes.Buffer)
		template.Index(ExampleItems, buffer)
		w.Write(buffer.Bytes())
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
