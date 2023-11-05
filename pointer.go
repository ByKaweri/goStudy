package main

import "fmt"

func changer(a *int) {
	*a = *a * *a

}
func main() {
	a := 100
	changer(&a)
	fmt.Printf("a: %v\n", a)
}
