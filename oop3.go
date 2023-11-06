package main

import "fmt"

type AnmailIF interface {
	Sleep()
	Eat()
	Color()
}

type Cat struct {
	color string
}

func (this *Cat) Eat() {
	fmt.Printf("\"cat eat\": %v\n", "cat eat")
}
func (thi *Cat) Sleep() {
	fmt.Printf("\"cat sleep\": %v\n", "cat sleep")
}
func (this *Cat) Color() {
	fmt.Printf("this.color: %v\n", this.color)
}

type Dog struct {
	color string
}

func (this *Dog) Eat() {
	fmt.Printf("\"cat eat\": %v\n", "cat eat")
}
func (thi *Dog) Sleep() {
	fmt.Printf("\"cat sleep\": %v\n", "cat sleep")
}
func (this *Dog) Color() {
	fmt.Printf("this.color: %v\n", this.color)
}
func main() {
	var ai AnmailIF = &Cat{color: "red"}
	ai.Color()
}
