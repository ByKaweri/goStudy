package main

import (
	"fmt"
	"os"
)

func main() {
	i := add(9, 10)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("getArgs(): %v\n", getArgs())
}
func add(num1 int, num2 int) int {
	return num1 + num2
}
func getArgs() (str string) {
	str = os.Args[0]
	return
}
