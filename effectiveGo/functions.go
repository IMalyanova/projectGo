package main

import "fmt"

func main() {

	var (
		x int
		b []byte
	)

	for i := 0; i < len(b);  {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}

func (file *File) Write(b []byte) (n int, err error)
//==================================================

func  nextInt(b []byte, i int)  (int, int) {
	for  ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for  ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return  x, i
}

//===================================================

