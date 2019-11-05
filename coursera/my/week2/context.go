package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	ctx, finish := context.WithCancel(context.Background())
	ressult := make(chan int, 1)

	for i := 0; i <= 10 ; i++  {
		go worker(ctx, i, ressult)
	}

	foundBy := <- ressult
	fmt.Println("ressult found by", foundBy)
	finish()

	time.Sleep(time.Second)
}

func  worker(ctx context.Context, workerNum int, out chan <- int)  {

	waitTime := time.Duration(rand.Intn(100) + 10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)

	select {
	case <- ctx.Done():
		return
	case <- time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}