package main

import "fmt"

func newTask() {
	for i := 0; i < 20; i++ {
		fmt.Println("new goroutine", i)
	}
}
func main() {
	go newTask()
	for i := 0; i < 100; i++ {
		fmt.Println("main goroutine", i)
	}
	fmt.Printf("\"main task exit\": %v\n", "main task exit")

}
