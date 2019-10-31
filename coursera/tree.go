package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var space string

func main() {


	//in := bufio.NewScanner(os.Stdin)
	//
	//for in.Scan() {
	//
	//	txt := in.Text()
	//	file, err := os.Create(txt)
	//
	//	if err != nil {
	//		// здесь перехватывается ошибка
	//		return
	//	}
	//
	//	fmt.Println(file)
	//	//separator := filepath.Separator



		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

			if info.IsDir(){
				space +=  "              "
				fmt.Println(space, info.Name())
			} else {
				space +=  "              "
				fmt.Println(space, info.Name(), "----", info.Size(), " byte")
				space = strings.Replace(space, "              ", "", 1)
			}
			return nil
		})
}
