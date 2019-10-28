package main

import (
	"bytes"
	"go/types"
	"sync"
)

type SyncedBuffer struct {
	lock sync.Mutex
	buffer bytes.Buffer
}

func main() {

	p := new(SyncedBuffer)
	var v SyncedBuffer
	//===========================
	
}
//===============================

//func NewFile(fd int, name string) *File {
//
//	if fd < 0 {
//		return  nil
//	}
//
//	f := new(File)
//	f.fd = fd
//	f.name = name
//	f.dirinfo = nil
//	f.nepipe = 0
//
//	return f
//}

												//тоже самое ниже, но только сокращено

func  NewFile (fd int, name string)*File  {
	if fd < 0 {
		return  nil
	}
	//f := File{fd, name, nil, 0}
	//return &f

	return &File {fd, name, nil, 0}             // тоже самое, что и выше две строки
	//return & File {fd: fd, name: name}  - можно ит ак, тоже самое
}

//===============================================
var p * [] int = new([] int)
var v [] int = make([] int, 100)

var p * [] int = new([] int)
*p = make ([] int, 100, 100)

v := make([] int, 100)

//===================================================

var p *[]int = new([]int)
var v []int = make([]int, 100)

var p *[]int = new([]int)
*p = make([]int, 100, 100)

v := make([]int, 100)
//============================================

func (f* File) Read (buf[] byte)(n int, err error)
n, err := f.Read( buf[0:32])


var n int
var err error
for i := 0;  i < 32; i ++ {
	nbytes, e := f.Read( buf[ i : i + 1])
	n += nbytes
	if nbytes == 0 || e != nil {
		err = e
		break
	}
}

//==============================================

func Append( slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l + len(data)) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[ 0: l + len(data)]
	copy( slice[l:], data)
	return  slice
}

