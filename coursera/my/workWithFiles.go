package my


//http://golang-book.ru/chapter-13-core-packages.html

// gofmt -w project1  - форматирует программу

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// создаем файл
	file, err := os.Create("test.txt")

	if err != nil {
		// здесь перехватывается ошибка
		return
	}
	defer file.Close()

	// запись в файл
	file.WriteString("test")



	// получить размер файла
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat)



	// получаем адреса файлов
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})


	//PathSeparator  - разделитель пути для ОС
}