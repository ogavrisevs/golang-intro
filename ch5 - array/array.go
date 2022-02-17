package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	simpleArry()
	autoSizeArry()
	arryRange()
}

func simpleArry() {
	fmt.Println("simpleArry")
	var x [5]int
	x[0] = 01
	x[4] = 5

	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))
}

func autoSizeArry() {
	fmt.Println("autoSizeArry")
	x := [...]int{1, 2, 3}
	x[0] = 11

	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))
	// x[3] = 3 -> invalid array index 3 (out of bounds for 3-element array)

}

func arryRange() {
	fmt.Println("range")
	x := [6]int{1, 2, 3, 4, 5, 6}

	for i, v := range x {
		fmt.Print(" this is elemnt : ", i, " with val : ", v)
	}

	for _, v := range x {
		fmt.Print(v)
	}
	fmt.Println()
}
