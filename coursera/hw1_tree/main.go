package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)

	if err != nil {
		panic(err.Error())
	}


	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		txt := in.Text()
	}
}

func dirTree (out *File, path string, printFiles) err error {

	file, err := os.NewFile()

	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return
	}

	fmt.Println(stat)
	return err
}