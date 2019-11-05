package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTimer(time.Second)

	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)

		if i >= 5 {
			//необходимо останавливать, иначе потечет
			ticker.Stop()
			break
		}
	}

	fmt.Println("total", i)
	//return


//	не может быть остановлен и собран сборщиком
//использовать можно только если должен работать вечно
c := time.Tick(time.Second)
i = 0

	for tickTime := range c{
		i++
		fmt.Println("sstep", i, "time", tickTime)
		if i >= 5 {
			break
		}

	}
}
