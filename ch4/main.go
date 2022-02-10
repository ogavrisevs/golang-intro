package main

import (
	"fmt"
)

func switchF(i int) {
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Dont knwo this umber")
	}
}

func main() {
	fmt.Println("Start...")

	for i := 0; i < 10; i++ {
		fmt.Println(" i : ", i)
		if i == 5 {
			fmt.Println(" i == 5")
		}
		switchF(i)
	}

}
