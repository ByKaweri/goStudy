package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age  int
	sex  bool
}

// 出现*User时无法反射获得参数
func (this User) Call() {
	fmt.Printf("\"Call\": %v\n", "Call")
}
func (this User) Tall() {
	fmt.Printf("\"tall\": %v\n", "tall")
}
func DoFileAndMethod(input interface{}) {
	fmt.Printf("reflect.TypeOf(input): %v\n", reflect.TypeOf(input))
	fmt.Printf("reflect.ValueOf(input): %v\n", reflect.ValueOf(input))
	//字段数量
	for i := 0; i < reflect.TypeOf(input).NumField(); i++ {
		field := reflect.TypeOf(input).Field(i)
		value := reflect.ValueOf(input).Interface()
		fmt.Println(field, value)
	}
	//方法数量
	for j := 0; j < reflect.TypeOf(input).NumMethod(); j++ {
		method := reflect.TypeOf(input).Method(j)
		fmt.Println("method", method)
	}
	fmt.Printf("reflect.TypeOf(input).NumMethod(): %v\n", reflect.TypeOf(input).NumMethod())

}
func main() {
	u := User{name: "kaweri", age: 18, sex: true}
	DoFileAndMethod(u)
}
