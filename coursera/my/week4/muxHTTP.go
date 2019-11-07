package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// HandleFunc вызывает базовый глобальный Multiplexor, который
	// обрабатывает запросы
	// - это функция обертки, которая предоставляет нам очень
	// быстрый интерфейс для внутренних объектов

	//Если нужно поднять несколько разных урлов или несколько разных серверов
	//Для этого:  Создаем объект с таким мультиплексором (http.NewServeMux()
	// и регистрируем у него нужные обработчики

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler2)

	server := http.Server {
		Addr:             ":8080",
		Handler:          mux,
		ReadTimeout:      10 * time.Second,
		WriteTimeout:     10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}


func handler2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "URL", r.URL.String())
}
