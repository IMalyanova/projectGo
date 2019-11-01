package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)
var space string


func main() {

	tree(".", space)
}



func tree (pathIn string, space string) {

	var arrayPath [] os.FileInfo
	var pathElement string

	sort.Slice(arrayPath, func(i, j int) bool {
		return false
	})

	filepath.Walk(pathIn, func(path string, info os.FileInfo, err error) error {
		arrayPath = append(arrayPath, info)
		pathElement = path

		return nil
	})

	for _, infoElement := range arrayPath {

		if !infoElement.IsDir() {
			//space +=  "              "
			fmt.Println(space, infoElement.Name(), "----", infoElement.Size(), " byte")
		} else {
			space +=  "              "
			fmt.Println(space, infoElement.Name())
			//gg(pathElement, space, infoElement.Name())
			//space = strings.Replace(space, "              ", "", 1)
		}
	}


}
//
//func gg( pathElement string, space string, name string ){
//
//	space +=  "              "
//	fmt.Println(space, name)
//	tree(pathElement, space)
//}




