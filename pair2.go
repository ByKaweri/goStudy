package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	C, err := os.OpenFile("../goStudy", os.O_RDWR, 0)
	fmt.Printf("C: %v\n", C)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var r io.Reader
	r = C
	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("HELLO T"))
}
