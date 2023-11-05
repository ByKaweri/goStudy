package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go")
	time.Sleep(time.Second)
	for i := 0; i <= 100; i++ {
		go fmt.Println("线程", i)
	}
}
