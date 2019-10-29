package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io"

	//=======

	_ "net / http / pprof"
)

func main() {

	m, ok := val.(json.Marshaler)

	//=======

	if _, ok := val.(json.Marshaler); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
	}

	//=======
	var _json.Marshaler = (*RawMessage)(nil)


	//=======

	fd, err := os.Open("test.go")

	if err != nil {
		log.Fatal(err)
	}

	//=======

	var _ = fmt.Printf
	var _io.Reader

	fd, err := os.Open("test.go")

	if err != nil {
		log.Fatal(err)
	}
	_ = fd

	//=======

	if _, err := os.Stat(path); os.IsExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
}
