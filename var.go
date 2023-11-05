package main

import "fmt"

func main() {
	var a int
	fmt.Printf("a: %v\n", a)

	var b int = 100
	fmt.Printf("b: %v\n", b)

	var c = 100
	fmt.Printf("c: %v\n", c)
	//只能在函数体内声明变量
	d := 100
	fmt.Printf("d: %v\n", d)

	var aa, bb = 100, 200
	fmt.Printf("aa: %v\n", aa)
	fmt.Printf("bb: %v\n", bb)
	var (
		cc      = 1
		dd      = 2
		ee      = "string"
		jj bool = true
	)
	fmt.Printf("jj: %v\n", jj)
	fmt.Printf("cc: %v\n", cc)
	fmt.Printf("dd: %v\n", dd)
	fmt.Printf("ee: %v\n", ee)
}
