package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriter struct {
	*Reader
	*Writer
}

//=======

type ReadWriter struct {
	reader *Reader
	writer *Writer
}

//=======

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

//=======

type Job struct {
	Command string
	*log.Logger
}

//=======

job.Println("starting now...")

//=======

func NewJob(command string, logger *log.Logger) *Job  {
	return &Job{command, logger}
}

//=======

job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}

//=======

func (job *Job) Printf(format string, args ...interface{})  {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}