package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
var space string

func main() {


}


func tree(path string, space string) {

	var arrayPath [] os.FileInfo

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		arrayPath = append(arrayPath, info)
		return nil
	})

	for _, infoElement := range arrayPath {

		if !infoElement.IsDir() {
			fmt.Println(space, infoElement.Name(), "----", infoElement.Size(), " byte")
		} else {
			fmt.Println(space, infoElement.Name())
		}
		setSpace(infoElement, space)

	}
}


func setSpace(infoElement os.FileInfo, space string) string {

	if infoElement.IsDir(){
		space +=  "              "
		//tree(infoElement, space)
		space = strings.Replace(space, "              ", "", 1)
	}

	return space
}



