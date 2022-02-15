package main

import "fmt"

func main() {
	fmt.Println("GO")
	intArry := []int{1, 2, 3, 4}
	fmt.Println(getSum(intArry))
	fmt.Println(returnTypeName())
	intVal, strVal := returnMultiVals()
	fmt.Println("returnMultiVals() intVal : ", intVal, " strVal : ", strVal)
	fmt.Println(variadicFunc(intArry...))
}

func getSum(intArry []int) int {
	fmt.Println("call getSum")
	var out int
	for _, v := range intArry {
		out += v
	}
	return out
}

func returnTypeName() (returnInt int) {
	fmt.Println("call returnTypeName")
	returnInt = 5
	return
}

func returnMultiVals() (returnInt int, returnStr string) {
	fmt.Println("call returnMultiVals")
	returnInt = 6
	returnStr = "from func returnMultiVals"
	return
}

func variadicFunc(intArry ...int) int {
	fmt.Println("call variadicFunc")
	var sumInt int = 1
	for _, v := range intArry {
		sumInt += v
	}
	return sumInt
}
