package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Fatal(err)
		}
		go sendNewMsgNotifications(ws)
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}



func sendNewMsgNotifications(client *websocket.Conn) {

	ticker := time.NewTicker(3 * time.Second)

	for{
		w, err := client.NextWriter(websocket.TextMessage)

		if err != nil {
			ticker.Stop()
			break
		}

		msg := newMesage()
		w.Write(msg)
		w.Close()

		<- ticker.C
	}
}

