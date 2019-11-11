package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperastions int32 = 0

func inc()  {
	atomic.AddInt32(&totalOperastions, 1) // атомарно
}

func main() {
	for i :=0; i < 1000; i++  {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation = ", totalOperastions)
}
