package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	arrSplit := strings.Split(s, " ")
	m := make(map[string]int)
	count:= 1

	for _, v := range arrSplit {
		// If m contain v count++
		_, ok := m[v]
		if ok {
			m[v] = count + 1
		}else {
			m[v] = count
		}

	}
	fmt.Println(m)
	return m
}

func main() {
	WordCount("Hi Hi my friend")
}