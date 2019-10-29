package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

type Stringer interface {
	String() string
}
var value interface{}

func main() {
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	//===========================

	value.(typeName)
	//=================================
	str := value.(string)
	//=================================

	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
	//=================================

	if str, ok := value.(string); ok {
		return str
	}else if str, ok := value.(Stringer); ok {
		return str.String()
	}

}
//====================================
type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less() (i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return  append(copy, s...)
}

func(s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	str := "["

	for i; elem := range s  {
		if i > 0 {
			str += " "
		}
	str += fmt.Sprint(elem)
	}
	return str + "]"
}
//==Conversions===================

func  (s Sequence) String() string  {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}
//=================================

func  (s Sequence) String() string  {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

//===Generality==============================

type Block interface {
	BlockSize() int
	Encrypt(dst, src []byte)
	Decrypt(dst, src []byte)
}

type Stream interface {
	XORKeyStream(dst, src []byte)
}
//======================================

func NewCTR(block Block, iv []byte) Stream
//========================================

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

//Simple counter server.===============
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP (w http.ResponseWriter, req *http.Request)  {
	ctr.n ++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

//=======================================

import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)

//=======================================

type Counter int

func (ctr *Counter) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}
//=====================================

type Chan chan *http.Request

func (ch Chan) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprintf(w, "notification sent")
}

//=====================================

func ArgServer (){
	fmt.Println(os.Args)
}

type HandlerFunc func(ResponceWriter, req *Request){
	f(w, req)
}
//=====================================

func ArgServer(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, os.Args)
}

//=====================================

http.Handle("/args", http.HandlerFunc(ArgServer))

//=====================================
