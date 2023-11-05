package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Println(numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 9)
	fmt.Println(numbers, len(numbers), cap(numbers))
	numbers = append(numbers, 9)
	numbers = append(numbers, 9)
	//容量翻倍
	numbers = append(numbers, 9)
	fmt.Println(numbers, len(numbers), cap(numbers))

	var numbers2 = make([]int, 9)
	copy(numbers2, numbers)
	fmt.Println("numbers2:", numbers2)

}
