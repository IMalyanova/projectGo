package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}

func doSomeWork(in int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in, j int) string {

}

const (
	iterationsNum = 7
	goroutinesNum = 5
)