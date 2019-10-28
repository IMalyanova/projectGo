package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var (
		x int
		b []byte
	)

	for i := 0; i < len(b);  {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
	//=======================
	b()
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

func ReadFull(r Reader, buf []byte) (n int, err error){
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

//===================================================

func Contents(filename string) (string, error){
	f, err := os.Open(filename)
	if err != nil{
		return  ", err"
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil{
			if err == io.EOF{
				break
			}
			return "", err
		}
	}
	return string(result), nil
}
//========================================
for i := 0; i < 5; i ++ {
	defer fmt.PrintF("%d ", i)
}
//========================================
func a(){
	trace("a")
	defer untrace("a")
}
//=========================================
func trace(s string) string{
	fmt.Println("entering: ", s)
	return s
}

func un(s string)  {
	fmt.Println("leaving: ", s)

}

func  a()  {
	defer  un( trace("a"))
	fmt.Println("in a")
}

func  b()  {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

