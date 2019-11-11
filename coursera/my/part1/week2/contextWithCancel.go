package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	workTime := 50 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	result := make(chan int, 1)

	for i := 0; i <= 10 ; i++  {
		go worker2(ctx, i, result)
	}
}

func worker2(ctx context.Context, workerNum int, out chan <- int)  {
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
