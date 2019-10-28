package main

import (
	"bytes"
	"sync"
)

func main() {

	type SyncedBuffer struct {
		lock sync.Mutex
		buffer bytes.Buffer
	}
	
}
