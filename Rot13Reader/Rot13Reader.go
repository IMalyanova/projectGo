package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	// reads from the io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetic characters.

	//if err == nil {
		n, err = rot13.reader.Read(p)
		var str string
		for i, v := range p {
			p[i] = v + 13
			str += string(p[i])

			fmt.Printf("n = %v, err = %v, p = %v\n", n, err, p)
	//	}
			fmt.Printf("%q\n", str)
			if err == io.EOF {
				break
			}
	}
	return n, err
}

func main() {
	//b := make([]byte, 8)
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	//r.Read(b)
	//io.Copy(os.Stdout, &r)
	io.CopyN(os.Stdout, &r, 20)
}
