package main

import "fmt"

func main() {
	cityMap := map[string]string{
		"China": "beijing",
		"USA":   "NewYork",
		"UK":    "Englen",
	}
	addCity(cityMap, "JP", "TokYO")
	printMap(cityMap)
	delCity(cityMap, "UK")
	printMap(cityMap)
}
func addCity(cityMap map[string]string, counter, city string) {
	cityMap[counter] = city
}
func delCity(cityMap map[string]string, counter string) {
	delete(cityMap, counter)
}
func printMap(Map map[string]string) {
	fmt.Printf("Map: %v\n", Map)
}
