package main

import "fmt"

var (
	Version = ""
	Branch = ""
)

func main(){
	fmt.Println("[start] starting version ", Version, Branch)
}