package main

import "fmt"

func main() {

	closureFunc := func(strIn string) (strOut string) {
		strOut = strIn + "_"
		return
	}

	fmt.Println(closureFunc("Hi its me"))

	obj := funcAsObject(" Start ")
	fmt.Println(obj())
	fmt.Println(obj())

}

func funcAsObject(strIn string) func() string {

	var localState string = strIn

	return func() (ret string) {
		localState += " call "
		return localState
	}

}
