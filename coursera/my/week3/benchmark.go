package main

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func main() {
	
}

func BenchmarkCodegen(b *testing.B) {

	for i := 0; i < b.N; i++ {
		u := &User{}
		u.Unpack(data)
	}
}

func BenchmarkReflect(b *testing.B) {

	for i := 0; i < b.N; i++ {
		u := &User{}
		UnpackReflect(u, data)
	}
}

func (in *User) Unpack(data []byte) error {
	r := bytes.NewReader(data)

	//ID
	var IDRaw uint32
	binary.Read(r, binary.LittleEndian, &IDRaw)
	in.ID = int(IDRaw)
}