package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 5, 6}
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	slice2 := make([]int, 8)
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
	var slice3 []int
	if slice3 == nil {
		println("null")
	}
}
