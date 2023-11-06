package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("a defer")
		func() {
			defer fmt.Println("b defer")
			runtime.Goexit()
			fmt.Printf("\"b\": %v\n", "b")
		}()
		fmt.Printf("\"a\": %v\n", "a")
	}()
	go func(a, b int) bool {
		fmt.Println(a, b)
		//返回值无法被接受，需使用channal
		return true
	}(1, 4)
	for {
	}
}
