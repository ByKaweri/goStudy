package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"myName"`
	sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		s := t.Field(i).Tag.Get("doc")
		fmt.Println(taginfo, s)
	}
}
func main() {
	r := resume{Name: "kaweri", sex: "男"}
	findTag(&r)
}
