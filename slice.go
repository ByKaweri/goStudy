package main

import "fmt"

func main() {
	//固定数组
	var arry [10]int
	arry2 := [3]int{1, 3}
	for i := 0; i < 10; i++ {
		arry[i] = i
	}
	fmt.Printf("arry2: %v\n", arry2)
	fmt.Printf("arry: %v\n", arry)

	//切片
	arr2 := []int{1, 5, 76}
	fmt.Printf("arr2: %v\n", arr2)
}

//值传递
func fixed(arr [4]int) {

}

//指针传递
func slice(arr []int) {

}
