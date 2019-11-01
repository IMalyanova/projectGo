package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var arrayPath [] os.FileInfo
	var space string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		arrayPath = append(arrayPath, info)
		return nil
	})

	for _, infoElement := range arrayPath {

		if !infoElement.IsDir(){
			fmt.Println(space, infoElement.Name(), "----", infoElement.Size(), " byte")
		} else {
			fmt.Println(space, infoElement.Name())
		}

	}

}

func jj(space string) string{



	return space
}