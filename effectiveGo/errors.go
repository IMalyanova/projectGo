package main

import (
	"fmt"
	"os"
)

func main() {

}

type error interface {
	Error() string
}

//=======

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

//=======

open /etc/passwx: no such file or directory

//=======

for try := 0; try < 2; try++ {
	file, err = os.Create(filename)

	if err == nil {
		return
	}

	if e, ok := errr.(*os.PathError); ok && e.Err == syscall.ENOSPC {
		deleteTempFiles()
		continue
	}

	return
}

//=======Panic

func CubeRoot(x float64) float64  {
	z := x/3

	for i := 0; i < 1e6; i ++  {
		prevz := z
		z -= (z * z * z - x) / (3 * z * z)

		if veryClose(z, prevz) {
			return z
		}
	}

	panic( fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

//=======

var user = os.Getenv("USER")

func init() {

	if usser == "" {
		panic("No value for $USER")
	}
}

//=======