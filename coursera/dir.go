package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main1() {

	//files, err := ioutil.ReadDir(".")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, f := range files {
	//	fmt.Println(f.Name())
	//}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
