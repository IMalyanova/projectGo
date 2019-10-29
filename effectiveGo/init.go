package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init")
}

// init function=====================
func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath may be overridden by --gopath flag on commsnd line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

//==================================
