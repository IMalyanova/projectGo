package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	go list.Sort()
}
//=======

func Announce(message string, delay time.Duration){
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}

//=======

ci := make(chan int)
cj := make(chan int, 0)
cs := make(chan *os.File, 100)

//=======

c := make(chan int)

go func(){
	list.Sort()
	c <- 1
}()

doSomethingForAWhile()
<- c

//=======

var sem = make(chan int, MaxOutstanding)

func handle(r *Request){
	sem <- 1
	process(r)
	<- sem
}

func Serve(queue chan *Request){

	for{
		req := <- queue
		go handle(req)
	}
}

//=======

func Serve(queue chan *Request){

	for req := range queue {
		sem <- 1
		go func() {
			process(req)
			<- sem
		}()
	}
}

//=======

func Serve(queue chan *Request){
	for req := range queue {
		req := req
		sem <-1

		go func() {
			process(req)
			<- sem
		}()
	}
}

//=======

req := req

//=======

func  handle(queue chan *Request)  {
	for r :=range queue {
		process(r)
	}
}

func Serve(clientRequests chan *Request, quit chan bool)  {
	for i := 0; i < MaxOutstanding; i ++ {
		go handle(clientRequests)
	}
	<- quit
}

//=======

type Request struct {
	args [] int
	f func([] int) int
	resultChan chan int
}

//=======

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

request := &Request {[]int{2, 4, 5}, sum, make(chan int)}
clientRequests <- request
fmt.Printf("answer: %d\n", <-request.resultChan)

//=======

func handle(queue chan *Request){
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

//=======Parallelization





