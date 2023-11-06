package main

import "fmt"

func fibonaci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case <-c:
			fmt.Println(x)
			x = y
			y = x + y

		case <-quit:
			fmt.Printf("\"exit\": %v\n", "exit")
			return
		}
	}
}
func main() {
	quit, c := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			c <- i
			fmt.Printf("main i: %v\n", i)
		}
		quit <- 0
	}()
	fibonaci(c, quit)
}
