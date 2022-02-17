package main

import "fmt"

func one() {
	fmt.Println("func One")
}

func two() {
	fmt.Println("func Two")
}

func main() {
	defer one()
	two()
}
