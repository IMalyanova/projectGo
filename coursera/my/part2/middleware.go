package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic", panicPage)

	// set middleware
	adminHandler := adminAuthMiddleware(adminMux)

	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.HandleFunc("/login", loginPage)
	siteMux.HandleFunc("/logout", logoutPage)
	siteMux.HandleFunc("/", mainPage)

	// set middleware
	siteHandler := accessLogMiddleware(siteMux)
	siteHandler = panicMiddleware(siteHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(" :8080", siteHandler)
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