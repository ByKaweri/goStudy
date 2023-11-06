package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		defer fmt.Println("goroutine over")
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("i: %v\n", i)
		}
	}()
	time.Sleep(1000)
	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Printf("num: %v\n", num)
	}
}
