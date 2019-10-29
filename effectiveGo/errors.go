package main

func main() {

}

type error interface {
	Error() string
}

//=======

type PathError struct {
	Op   string
	Path string
	Err  error
}
