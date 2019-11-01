package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
var space string = "              "


func main() {

	tree(".", space)
}



func tree (pathIn string, space string) {

	var arrayPath [] os.FileInfo
	var pathElement string

	filepath.Walk(pathIn, func(path string, info os.FileInfo, err error) error {
		arrayPath = append(arrayPath, info)
		pathElement = path

		return nil
	})

	for _, infoElement := range arrayPath {

		if !infoElement.IsDir() {
			fmt.Println(space, infoElement.Name(), "----", infoElement.Size(), " byte")
		} else {
			space = strings.Replace(space, "              ", "", 1)
			fmt.Println(space, infoElement.Name())
			space +=  "              "
			tree(pathElement, space)

		}
	}
}





