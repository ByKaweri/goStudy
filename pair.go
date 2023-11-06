package main

import "fmt"

func main() {
	//pair <type:string> <value:"kaweri">

	var a string = "kaweri"
	var allType interface{}
	//pair <type:string> <value:"kaweri">
	allType = a
	s, ok := allType.(string)
	if ok {
		fmt.Printf("s: %v\n", s)
	}

}
