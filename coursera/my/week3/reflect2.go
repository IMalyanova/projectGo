package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func main() {

	data byte := []{
		128, 36, 17, 0,

		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}
	u := new(User)
	err := UnpackReflect(u, data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", u)
}

type User struct {
	ID       int
	RealName string `unpack:"-"`
	Login    string
	Flags    int
}

func UnpackReflect(u interface{}, data []byte) error {

	r := bytes.NewReader(data)
	val := reflect.ValueOf(u).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.NumField(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32
			binary.Read(r, binary.LittleEndian, &value)
			valueField.Set(reflect.ValueOf(int(value)))
		case reflect.String:
			var lenRaw uint32
			binary.Read(r, binary.LittleEndian, &lenRaw)



		}
	}
}
