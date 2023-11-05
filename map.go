package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 4)
	fmt.Printf("myMap: %v\n", myMap)
	myMap["one"] = "1"
	myMap["two"] = "2"
	fmt.Printf("myMap: %v\n", myMap)
	myMap2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("myMap2: %v\n", myMap2)
}
