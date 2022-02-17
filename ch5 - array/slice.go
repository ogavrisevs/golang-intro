package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	sliceSimple()
	sliceCap()
	sliceSlicing()
	sliceAppend()
	sliceCoppy()
}

func sliceSimple() {
	fmt.Println("sliceSimple")

	var arryOpen []int
	fmt.Println("len(arryOpen) : ", len(arryOpen))
	//arryOpen[0] = 100  -> panic: runtime error: index out of range [0] with length 0

	arrySize5 := make([]int, 5)
	arrySize5[0] = 10
	fmt.Println("len(arrySize5) : ", len(arrySize5))
	for _, v := range arrySize5 {
		fmt.Print(" ", v)
	}
	fmt.Println("")
}

func sliceCap() {
	fmt.Println("SliceCap")
	sliceSize1Cap5 := make([]int, 1, 5)
	fmt.Println("len(sliceSize1Cap5) : ", len(sliceSize1Cap5))
	sliceSize1Cap5[0] = 10
	fmt.Println("len(sliceSize1Cap5) : ", len(sliceSize1Cap5))
	fmt.Println("cap(sliceSize1Cap5) : ", cap(sliceSize1Cap5))
	fmt.Println("sliceSize1Cap5 : ", sliceSize1Cap5)
}

func sliceSlicing() {
	fmt.Println("sliceSlicing")

	stringSlice := []string{"aaaa", "bbbb", "cccc", "dddd", "eeee"}
	fmt.Println("stringSlice : ", stringSlice, "cap(stringSlice)", cap(stringSlice), "len(stringSlice)", len(stringSlice))
	fmt.Println("stringSlice[1:3]", stringSlice[1:3], "cap(stringSlice[1:3])", cap(stringSlice[1:3]), "len(stringSlice[1:3])", len(stringSlice[1:3]))
	fmt.Println("stringSlice[:2]", stringSlice[:2], "cap(stringSlice[:2])", cap(stringSlice[:2]), "len(stringSlice[:2])", len(stringSlice[:2]))
	fmt.Println("stringSlice[2:]", stringSlice[2:])
	fmt.Println("stringSlice : ", stringSlice)

	stringSlice = stringSlice[2:]
	stringSlice = stringSlice[:cap(stringSlice)]
	fmt.Println("stringSlice : ", stringSlice)
}

func sliceAppend() {
	fmt.Println("sliceAppend")

	slice1 := make([]int, 5)
	slice1[0] = 1000
	slice1[4] = 5000
	fmt.Println("len(slice1) : ", len(slice1))

	slice1 = append(slice1, 22, 33)
	fmt.Println("len(slice1) : ", len(slice1))

	slice1 = append(slice1, 44, 55)
	fmt.Println("len(slice1) : ", len(slice1))

	for i, v := range slice1 {
		fmt.Print("[", i, "]:", v, ", ")
	}
	fmt.Println(" ")
}

func sliceCoppy() {
	fmt.Println("sliceCoppy")

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
