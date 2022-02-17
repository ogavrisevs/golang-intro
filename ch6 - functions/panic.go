package main

import (
	"fmt"
)

func main() {

	defer func() {
		str := recover()
		fmt.Println(" will recover from : ", str)
	}()

	panic("panic from main")
}
