package main

import "fmt"

func main() {
	x := []int {1,2,3}
	x = append(x, 4,5,6,7)
	//================================

	x = [] int {1,2,3}
	y := [] int {40,50,60}
	x = append( y, 7,8,90)
	fmt.Println(x, y)
}
//func append(slice []T, elements ...T) []T
//=========================================

