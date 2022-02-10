package main

import "fmt"

var strVarGlob string = "globalVar "

func main() {

	var unIntVar uint = +100
	var intVar int = -100
	fmt.Println("this is int :", intVar)
	fmt.Println("this is uint :", unIntVar)

	var strVar string
	strVar = "myString"
	fmt.Println("this is string : ", strVar)
	fmt.Println("this is string.leng : ", len(strVar))

	var boolVar bool
	boolVar = true
	fmt.Println(boolVar)

	fmt.Println(strVarGlob)

	strVar += strVarGlob
	fmt.Println(strVar)

	const const1 string = "CONSTVAR"
	fmt.Println(const1)

	var (
		a = 1
		b = 11
		c = 111
	)
	fmt.Println("print : ", a, b, c)

}
