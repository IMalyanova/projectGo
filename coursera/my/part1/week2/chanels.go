package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)

	go func(in chan int) {
		val := <- in
		fmt.Println("Go: get from chan", val)
		fmt.Println("Go: after read from chan")
	}(ch1)

	ch1 <- 42
	ch1 <- 100500

	fmt.Println("MAIN: after put to chan")
	fmt.Scanln()
}
