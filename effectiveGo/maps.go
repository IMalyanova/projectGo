package main

import (
	"fmt"
	"log"
)

var timeZone = map[string]int{
	"UTC": 0*60*60,
	"EST": -5*60*60,
	"CST": -6*60*60,
	"MST": -7*60*60,
	"PST": -8,*60*60,
}

func main() {

	offset := timeZone["EST"]
	//================

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,

	}

	if attended [person]{
		fmt.Println(person, "Was at the meeting")
	}
	//=============================


	//==============================


}
func offset(tz string) int {

	var seconds int
	var ok bool
	seconds, ok = timeZone[tz]

	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println( "Unknown time zone: ", tz)
	return 0
}
//=====================================================

_, present := timeZone[tz]
delete(timeZone, "PDT")