package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	data := []byte(`some test`)
	str := bytesToStr(data)
	//str := string(data)
	fmt.Printf("type: %T, value: %+v\n", str, str)
}



func bytesToStr(data []byte) string {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	fmt.Printf("type: %T, value: %+v\n", h, h)
	fmt.Printf("type: %T, value:%+v\n", h.Data, h.Data)
	shdr := reflect.StringHeader{Data: h.Data, Len: h.Len}
	fmt.Prinf("type: %T, value: %+v\n", shdr, shdr)
	return *(*string)(unsafe.Pointer(&shdr))
}