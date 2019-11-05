package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	// необходиом использовать буферезированный канал
	result := make(chan string, 1)
	go func(out chan <- string ) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")
		out <- "32 comments"
	}(result)

	return result
}

func getPage() {
	resultCh := getComments()

	time.Sleep(1 * time.Second)
}
