package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	
}



func uploadRawBody(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	p := &Params{}
	err = json
}


