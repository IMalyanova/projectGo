package main
// fibonacci is a function that returns
// a function that returns an int.

import (
	"fmt"
	"math"
)

func fibonacci() (func(i int) int) {

	var SQRT5 = math.Sqrt(5)
	var PHI = (SQRT5 + 1) / 2

	return func(i int) int {
		return int(math.Pow(PHI, float64(i)) / SQRT5 + 0.5)
	}
}

func main() {
	count:= 20
	f := fibonacci()

	for i := 0; i < count; i++ {
		fmt.Println(f(i))
	}
}