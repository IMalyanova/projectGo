package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// пустая строка по-умолчанию
	var str string

	// со спец символами
	var hello string = "Привет\n\t"

	// без спец символов
	var world string = `Мир\n\t`

	fmt.Println("str", str)
	fmt.Println("hello", hello)
	fmt.Println("world", world)
//
//var str string
//	var hello string = "Привет\n\t"
//	var world string = 'Мир \n\t'
//
//	fmt.Println("str", str)
//	fmt.Println("hello", hello)
//	fmt.Println("world", world)
//
	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"
	hi := "你好，世界"

	fmt.Println("helloWorld", helloWorld)
	fmt.Println("hi", hi)

	// одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) для UTF-8 символов
	var someChinese rune = '茶'

	fmt.Println(rawBinary, someChinese)
//
//var helloWorld = "sdjf"
//hi := "ac dcds"
//
//
//fmt.Println("helloWorld", helloWorld)
//fmt.Println("hi", hi)
//
//var rawBinary byte = '\x27'
//var someChinese rune = '#$%'
//fmt.Println(rawBinary, someChinese)
//



	helloWorld = "Привет Мир"
	// конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"

	fmt.Println(helloWorld, andGoodMorning)

	// строки неизменяемы
	// cannot assign to helloWorld[0]
	// helloWorld[0] = 72

	// получение длины строки
	byteLen := len(helloWorld)                    // 19 байт
	symbols := utf8.RuneCountInString(helloWorld) // 10 рун

	fmt.Println(byteLen, symbols)

	// получение подстроки, в байтах, не символах!
	hello = helloWorld[:12] // Привет, 0-11 байты
	H := helloWorld[0]      // byte, 72, не "П"
	fmt.Println(H)

	//
	//andGoodMorning := helloWorld + " и доброе утро"
	//fmt.Println(helloWorld, andGoodMorning)
	//byteLen := len(helloWorld)
	//symbols := utf8.RuneCountInString(helloWorld)
	//fmt.Println(byteLen, symbols)
	//
	//hello = helloWorld[ :12]
	//H := helloWorld[0]
	//fmt.Println(H)
	//

	// конвертация в слайс байт и обратно
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(byteString, helloWorld)

	//
	//byteString := []byte(helloWorld)
	//helloWorld = string(byteString)
	// fmt.Println(byteString, helloWorld)
	//
}
