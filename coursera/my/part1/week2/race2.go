package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}

	for i:= 0; i < 5; i++  {
		go func(counters map[int]int, th int, mu *sync.Mutex) {

			for j :=0; j < 5; j++  {
				mu.Lock()
				counters[th * 10 + j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	fmt.Scanln("couters result", counters)
}
