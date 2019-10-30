package main

import (
	"fmt"
	"io"
	"strings"
	//"golang.org/x/tour/reader"
)

type MyReader struct{
	Ch int
}

func (r *MyReader) myRead() {
	b := make([]byte, 8)
	for ;; {
		reader := strings.NewReader(string(r.Ch))
		n, err := reader.Read(b)
		fmt.Printf("%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	mr := &MyReader{65}
	//reader.Validate(mr)
	mr.myRead()
}
