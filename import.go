package main

import (
	"fmt"
	_ "fmt"     //匿名导入 无法使用
	. "os"      //直接使用方法名 无需os.
	system "os" //别名
)

func main() {
	fmt.Printf("system.Args[0]: %v\n", system.Args[0])
	fmt.Printf("Args[0]: %v\n", Args[0])
}
