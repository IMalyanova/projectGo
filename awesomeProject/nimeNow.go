package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s []uint8

	x := uint8(dx)
	s = append(s, x)
	pow := make([][]uint8, dy)

	for i:= 0; i < dy; i++ {
		pow[i] = s
		//fmt.Println(pow[i])
	}
	return pow
}

func main() {

	pic.Show(Pic(5,7))
}
