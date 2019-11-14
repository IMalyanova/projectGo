package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	seessManager := NewSessManager()

	rpc.Register(sessManager)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":8081")
	if e != nill {
		log.Fatal("listen error:", e)
	}

	fmt.Println("starting server at :8081")
	http.Serve(l, nil)
}
