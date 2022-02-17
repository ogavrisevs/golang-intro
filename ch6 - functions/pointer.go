package main

import "fmt"

func main() {
	var i1 int = 11
	fmt.Println(i1)

	i2 := 22
	var intPointer *int
	intPointer = &i2
	fmt.Println(*intPointer)
	*intPointer = 33
	fmt.Println(*intPointer)

	var intPointer1 *int
	var intPointer2 *int
	intVal := 333
	intPointer1 = &intVal
	intPointer2 = &intVal
	fmt.Println("two pointer one val", *intPointer1, *intPointer2)
	intVal = 444
	fmt.Println("two pointer one val", *intPointer1, *intPointer2)

}
