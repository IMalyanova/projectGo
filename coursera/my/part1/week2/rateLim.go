package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)

	for i := 0; i < goroutinesNum; i++  {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func starWorker (in int, wg *sync.WaitGroup, quotaCh chan struct{}){

	quotaCh <- struct{}{}
	defer wg.Done()

	for j :=0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		// if j%2 == 0{
		//<- quotaCh
		//quotaCh <- struct{}{}
		//}

		runtime.Gosched()
	}
	<- quotaCh // ratelim.go, возвращаем слот

}
