package main

import (
	"fmt"
	"reflect"
)

func main() {
	
}

func  PrintReflect(u interface{}) error  {
	val := reflect.ValueOf(u).Elem()
	
	fmt.Printf("%T have %d fields:\n", u, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		
		fmt.Printf("\tname=%v, type=%v, value=%v, tag=`%v` \n",
			typeField.Name,
			typeField.Type.Kind(),
			valueField,
			typeField.Tag, 			
			)
	}
	return nil
}

func  main ()  {

	u := &User {
		ID:     42,
		RealName:  "rvasily",
	}
}