package main

import (
	"fmt"
	"os"
)

func main() {


	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	//=========================================
	var x uint64 = 1 << 64 - 1
	fmt.Printf("% d% x; % d% x \n", x, x, int64 (x),  int64(x))
	//=========================================
//	fmt.Printf("%v \n", timeZone)
	//=========================================

	type T struct {
		a int
		b float64
		c string
	}

	t := &T {7, -2.35, "abc \tdef"}
	fmt.Printf("% v \n", t)
	fmt.Printf("% + v \n", t)
	fmt.Printf("% #v \n", t)
//	fmt.Printf("%#v \n", timeZone)
	//================================================
	fmt.Printf("%v\n", t)
	//================================================
}
	//================================================

func (t *T) String() string {
	return  fmt.Sprintf("%d %g %q", t.a, t.b, t.c)
}


//error if %s in String:
type MyString string

func (m MyString) String() string  {
	return  fmt.Sprintf("MyString =% s", m)// Ошибка: будет повторяться вечно.
}
// исправить ошибку можно самим преобразовав в String
// , чтобы не вызывался метод String повторно: string(m)
//====================================================

func Printf (format string, v ...interface{}) (n int, err error){}
//=====================================================
func Println(v ...interface{})  {
	std.Output(2, fmt.Sprintln(v...))
}
//==================================================
func  Min(a ...int) int  {
	min := int(^uint(0) >> 1)
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
//=====================================================
