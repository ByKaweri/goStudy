package main

import "fmt"

func main() {
	for i := 0; i < 8; i++ {
		defer fmt.Printf("defer i: %v\n", i)
		fmt.Printf("i: %v\n", i)
	}
}
