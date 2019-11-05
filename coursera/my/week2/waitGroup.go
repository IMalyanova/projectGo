package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	for i := 0; i < goroutinesNum; i++  {
		wg.Add(1)
		go startWorker(i, wg)
	}

	time.Sleep(time.Millisecond)

	//fmt.Scanln
	wg.Wait()
}



func startWorker(in int, wg *sync.WaitGroup)  {

	defer  wg.Done()

	for j := 0; j < iterationsNum; j++  {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}