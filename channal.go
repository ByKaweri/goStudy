package main

import "fmt"

func main() {
	c := make(chan int, 8)
	go func() {
		defer fmt.Println("defere")
		fmt.Printf("\"a\": %v\n", "a")
		c <- 111
	}()
	str := <-c
	fmt.Printf("str: %v\n", str)
}
