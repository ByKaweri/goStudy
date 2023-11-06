package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (this *Human) Eat() {
	fmt.Printf("this.name: %v Eat\n", this.name)
}

type SuperMan struct {
	Human //继承Human
	level int
}

func (this *SuperMan) Fly() {
	fmt.Printf("this.name: %v Fly\n", this.name)
}
func main() {
	h := Human{name: "xiaoli", age: 10}
	fmt.Printf("h: %v\n", h)
	sm := SuperMan{Human{name: "Kalak", age: 18}, 88}
	fmt.Printf("sm: %v\n", sm)
	sm.Fly()
}
