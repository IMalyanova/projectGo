package main

import "fmt"

type MyStruct struct {
	userID 	int
	DataJSON []byte
}


// test
func TestError(isOK bool) error {
	if !isOK {
		//fmt.Errorf("failed")
		return fmt.Errorf("failed")
	}
	return nil
}

//test example
func Test() {
	flag := true
	result := TestError(flag)
	//if err != nil
	if result != nil {
		println(result)
	}

	//fmt.Printf("result is\n", result)
	//fmt.Printf("%v is %v", flag)
	fmt.Printf("result is %v\n", result)
	fmt.Printf("%v is %v", flag, result)
}