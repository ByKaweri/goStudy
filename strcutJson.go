package main

import (
	"encoding/json"
	"fmt"
)

type Movice struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	m := Movice{Title: "2019", Year: 2002, Price: 9, Actors: []string{"JK", "KI"}}
	jsonStr, err := json.Marshal(m)
	if err == nil {
		fmt.Printf("jsonStr: %s\n", jsonStr)
	}
	myMovice := Movice{}
	err2 := json.Unmarshal(jsonStr, &myMovice)
	if err2 == nil {
		fmt.Printf("myMovice: %v\n", myMovice)
	}
}
