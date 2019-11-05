package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func main() {

	tree(".", "")
}



func tree (pathElement string, space string) {

	var space2 string
	files, err := ioutil.ReadDir(pathElement)
	
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		space := space + "│"
		space2 = strings.TrimRight(space, "│")
		space2 += "├───"

		if !file.IsDir() {
			fmt.Println(space2, file.Name(), " (", file.Size(), "b)")
		} else {
			fmt.Println(space2, file.Name())
			space += "\t"
			tree( pathElement + string(os.PathSeparator) + file.Name() + string(os.PathSeparator), space)
			space = strings.Replace(space, "\t", "", 1)
		}
	}

}


