package main

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {

}

//============

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	*p = slice
}

//==============

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	*p = slice
	return len(data), nil
}

//===================

func main() {

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
}
