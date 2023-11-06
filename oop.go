package main

import "fmt"

//使用大小写区分是否可被外部访问
type Hero struct {
	name  string
	ad    int
	level int
}

//必须使用指针，否者只是值传递
func (this *Hero) GetName() string {
	return this.name
}
func (this *Hero) Show() {
	fmt.Printf("this: %v\n", this)
}
func main() {
	h := Hero{
		name:  "My",
		ad:    99,
		level: 88,
	}
	h.Show()
}
