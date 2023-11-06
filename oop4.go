package main

import "fmt"

//万能类型interface
func Myfunc(arg interface{}) {
	//判断interface的断言机制
	vakue, ok := arg.(string)
	if ok {
		fmt.Printf("vakue: %v\n", vakue)
	} else {
		fmt.Printf("\"not is string\": %v\n", "not is string")
	}
}
func main() {
	Myfunc("jjjk")
	Myfunc(1)
}
