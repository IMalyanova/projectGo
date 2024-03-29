package main

import "fmt"

var sessManager SessionManagerI


func main() {



	sessManager = NewSessManager()


	// создаем сессию
	sessId, err := sessManager.Create(
		&Session {
			Login:		"rvasily",
			Useragent:	"chrome",
		})
	fmt.Println("sessId", sessId, err)

	// проверяем сессию
	sess := sessManager.Check(
		&SessionID{
			ID: sessId.ID,
		})
	fmt.Println("sess", sess)

	//удаляем сессию
	sessManager.Delete(
		&SessionID{
			ID: sessId.ID,
		})

	// проверяем еще раз
	sessManager.Check(
		&SessionID{
			ID: sessId.ID,
		})

	fmt.Println("sess", sess)
}
