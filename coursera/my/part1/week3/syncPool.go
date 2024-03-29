package main

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

func main() {
	
}

const iterNum = 100

type PublicPage struct {

	ID     int
	Name   string
	Url    string
	ownerID int
	ImageUrl string
	Tags  [] string
	Description string
	Rules string
}

var CoolGolangPublic = PublicPage{
	ID:          1,
	Name:        "CoolGolangPublic",
	Url:         "http://example.com",
	ownerID:     100500,
	ImageUrl:    "http://example.com/img.png",
	Tags:        []string{"programming", "go", "golang"},
	Description: "Best page about golang programming",
	Rules:       "",
}

var Pages = []PublicPage{
	CoolGolangPublic,
	CoolGolangPublic,
	CoolGolangPublic,
}

func BenchmarkAllocNew(b *testing.B)  {
	b.RunParallel(func(pb *tesing.PB) {

		for pb.Next() {
			data := bytes.NewBuffer(make([]byte, 0, 64))
			_= json.NewEncoder(data).Encode(Pages)
		}
	})
}

var dataPool = sync.Pool {
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	},
}

func BenchmarkAllocPool(b *testing.B){
	b.RunParallel(func(pb *testing.PB){

		for pb.Next() {
			data := dataPool.Get().(*bytes.Buffer)
			_= json.NewEncoder(data).Encode(Pages)
			data.Reset()
			dataPool.Put(data)
		}
	})
}