package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	simpleMap()
	stringMap()
	string2dMap()
}

func simpleMap() {
	fmt.Println("simpleMap")
	x := make(map[string]int)
	x["one"] = 1
	fmt.Println(x)
}

func stringMap() {
	fmt.Println("stringMap")
	mapStr := make(map[string]string)
	mapStr["one"] = "one"
	mapStr["two"] = "two"
	mapStr["three"] = "three"
	fmt.Println("len(mapStr) : ", len(mapStr))
	fmt.Println(mapStr)
	for i, v := range mapStr {
		fmt.Println("[", i, "]:", v)
	}
}

func string2dMap() {
	fmt.Println("string2dMap")
	map2dStr := make(map[string]map[string]string)
	mapStr := make(map[string]string)
	mapStr["one"] = "one"
	mapStr["two"] = "two"
	mapStr["three"] = "three"

	map2dStr["firstEl"] = mapStr
	map2dStr["secondEl"] = map[string]string{"key1": "val1", "key2": "val2"}
	fmt.Println(map2dStr)

}
