package main

import "fmt"

func main() {
	c := make(chan int)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		c <- i
	// 	}
	// 	//不关闭可能造成死锁
	// 	close(c)
	// }()
	// for {
	// 	if num, ok := <-c; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }

	go func() {
		for j := 5; j < 10; j++ {
			c <- j
		}
		//不关闭可能造成死锁
		close(c)
	}()
	for e := range c {
		fmt.Printf("e: %v\n", e)
	}

}
