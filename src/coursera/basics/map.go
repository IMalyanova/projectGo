package main

import "fmt"

func main() {
	// инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}

	// сразу с нужной ёмкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)
//
//var user map[string]string = map[string]string {
//	"name": "Vasily",
//	"LastName": "Romanov",
//}
//
//profile := make(map[string]string, 10)
//mapLength := len(user)
//fmt.Println("#{mapLenght} #{profile}\n")
//
//var user map[string]string = map[string]string {
//	"name": "Vasily",
//	"LastName": "Romanov",
//}

//
	// если ключа нет - вернёт значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	// удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
////
//	mName := user["middleName"]
//	fmt.Println("mName:", mName)
//
//	mName, mNameExist := user["midddleName"]
//	fmt.Println("mName:", mName, "mNameExist:", mNameExist)
//
//	_, mNameExist2 := user["middleName"]
//	fmt.Println("mNameExist2", mNameExist2)
//
//	delete(user, "lastName")
//	fmt.Printf("#{user}\n")
////
