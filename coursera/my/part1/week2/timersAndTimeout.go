package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)

	select {
	case <- timer.C:
		fmt.Println("timer.C timeout happend")
	case <- time.After(time.Minute):
		fmt.Println("time.After timeout happend")
	case result := <- longSQLQuery():
		if !timer.Stop() {
			<- timer.C
		}
		fmt.Println("operation result:", result)
	}
}

func longSQLQuery() chan interface{} {
	ch := make(chan interface{})
	return ch
}
