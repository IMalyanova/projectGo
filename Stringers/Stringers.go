package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {

	//var array []string
	var str string
	var v byte

	for i := 0; i < len(ipAddr); i++ {
		v = ipAddr[i]
		if i != len(ipAddr) - 1{
			str += fmt.Sprintf("%v.", v)
		}else {
			str += fmt.Sprintf("%v", v)
		}
	}
	return fmt.Sprintf(str)
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
