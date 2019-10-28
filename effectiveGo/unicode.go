package main

import "fmt"

func main() {
	for pos, char := range "日本\x80語" {
		fmt.Fprintf("character %#U starts at byte position %d\n", string(char), pos)
	}
}
