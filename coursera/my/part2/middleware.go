package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	
}



func pageWithAllChecks(w http.ResponseWriter, r *http.Request) {

	defer func() {

		if err := recover(); err != nil {

			fmt.Println("recovered", err)
		}
	}()


	defer func(start time.Time) {

		fmt.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	}(time.Now())


	_, err := r.Cookie("session_id")
	// Учебный пример! Это не проверка авторизации!

	if err != nil {

		fmt.Println("no auth at", r.URL.Path)
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	// Your logic
}