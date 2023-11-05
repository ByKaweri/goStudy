package main

import "fmt"

const (
	//iota只能与const配合使用
	beijin = iota
	shanghai
	shenzheng
	guangzou
)
const (
	_ = iota
	a = 1 << iota
	b
	c
	d
)
const (
	j, k = iota, iota + 1
	u, i = iota, iota + 1
)

func main() {
	const length int = 10

	// length = 100
	// fmt.Printf("length: %v\n", length)
	fmt.Printf("beijin: %v\n", beijin)
	fmt.Printf("shanghai: %v\n", shanghai)
	fmt.Printf("shenzheng: %v\n", shenzheng)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("j: %v\n", j)
	fmt.Printf("k: %v\n", k)
	fmt.Printf("u: %v\n", u)
	fmt.Printf("i: %v\n", i)
}
