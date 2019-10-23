package main

import (
	"math/rand"
	"time"
	"golang.org/x/tour/pic"
)

func Pic(dx int, dy int) [][]uint8 {

	var array [][]uint8
	rand.Seed(time.Now().UnixNano())
	dx= randomInt(2, 12)
	dy= randomInt(3, 30)
	slice := array[0:dy]

	arx := [] uint8 {2, 3, 5, 7, 11, 13, 20, 24, 27, 29, 32, 35}

	for i:= 0; i < dy; i++ {
		slice[i] = arx[0:dx]
	}
	return slice
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	pic.Show(Pic(4,5))
}
