package main

import "time"

func main() {

	workerInput := make(chan string, 2)

	for i := 0; i < goroutinesNum; i++  {
		go startWorker(i, workerInput)
	}

	months := []string { "Январь", "Февраль", "Март", "Апрель",
		"Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь",
		"Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		workerInput <- monthName
	}
	close(workerInput) // попробуйте закоментировать

	time.Sleep(time.Millisecond)
}
