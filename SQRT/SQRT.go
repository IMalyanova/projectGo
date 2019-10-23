package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	return float64(x * x)
}

func main() {
	fmt.Println(Sqrt(2))
}
