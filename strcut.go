package main

import "fmt"

type myint int
type Book struct {
	title  string
	author string
}

//指针传递
func changeBook_(book *Book) {

}

//值传递
func changeBook(book Book) {

}
func main() {
	var b Book
	b.title = "goLang"
	b.author = "my"
	fmt.Printf("b: %v\n", b)
}
