package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type tplParams struct {
	URL     string
	Browser string
}

const EXAMPLE = `
	Browser {{.Bowser}}
	
	you at {{.URL}}`



func handle (w http.ResponseWriter, r *http.Request) {

	tmpl := template.New(`example`)
	tmpl, _ = tmpl.Parse(EXAMPLE)

	params := tplParams {
		URL:     r.URL.String(),
		Browser: r.UserAgent(),
	}

	tmpl.Execute(w, params)
}




func main() {

	http.HandleFunc("/", handle)

	fmt.Println("starting server at :8080")
}
