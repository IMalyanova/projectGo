package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func main() {

	out := new(bytes.Buffer)                                                                    //

	if !(len(os.Args) == 2 || len(os.Args) == 3) {                                      //
		panic("usage go run main.go . [-f]")                                        //
	}                                                                                     //

	path := os.Args[1]                                                                   //
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"                                //
	err := dirTree(out, path, printFiles)                                                //
                                                                                         //
	if err != nil {                                                                     //
		panic(err.Error())                                                              //
	}                                                                                  //


}




func dirTree (out *File, path string, printFiles bool) err error {

	var space string = "│"
	tree(".", space)


	if err != nil {
		return err
	}

	return err
}





func tree (pathElement string, space string) {

	files, err := ioutil.ReadDir(pathElement)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		space2 := space + "├───"
		if space2 == "│├───" {
			space2 = "├───"
		}

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
