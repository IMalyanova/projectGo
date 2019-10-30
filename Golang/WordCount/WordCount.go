package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	arrSplit := strings.Split(s, " ")
	m := make(map[string]int)
	count := 1

	for _, v := range arrSplit {
		v = strings.Trim(v, "!.,:?;")
		// If m contain v count++
		_, ok := m[v]
		if ok {
			m[v] = count + 1

		} else {
			m[v] = count
		}
	}
	return m
}

func main() {
	wc.Test(WordCount("Hi Hi my friend"))
}
