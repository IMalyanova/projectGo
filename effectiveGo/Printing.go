package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
}
