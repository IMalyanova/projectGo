package main

import "log"

func main() {

}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Work failed: ", err)
		}
	}()
	do(work)
}

//=======

type Error string

func (e Error) Error() string {
	return string(e)
}

func (regexp *Regexp) error(err string) {
	panic(Error(err))
}

func Compile(str string) (regexp *Regexp, err error) {

	regexp = new(Regexp)

	defer func() {

		if e := recover(); e != nil {
			regexp = nil
			err = e.(Error)
		}
	}()
	return regexp.doParse(str), nil
}

//https://golang.org/doc/effective_go.html

//=======

if pos == 0 {
	re.error (" '*' illegal at start of expression")
}

//=======

