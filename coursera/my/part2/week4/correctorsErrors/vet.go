package main

import "fmt"

func main() {
	flag := true
	result := testError(flag)
	//fmt.Printf("result is\n", result)
	fmt.Printf("result is %v\n", result)
	//fmt.Printf("%v is %v\n", flag)
	fmt.Printf("%v is %v\n", flag, result)
}



func testError(isOK bool) error {
	if !isOK {
		//fmt.Errorf("failed")
		return fmt.Errorf("failed")
	}
	return nil
}