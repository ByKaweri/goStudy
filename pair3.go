package main

import "fmt"

type Reader interface {
	ReadeBook()
}
type Writer interface {
	WriteBook()
}
type Book struct {
}

func (this *Book) ReadeBook() {
	fmt.Printf("\"read book\": %v\n", "read book")
}
func (this *Book) WriteBook() {
	fmt.Printf("\"Writer book\": %v\n", "Writer book")
}

func main() {
	b := &Book{}
	var r Reader
	r = b
	r.ReadeBook()
}
