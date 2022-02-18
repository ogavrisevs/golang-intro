package main

import "fmt"

type Characteristics interface {
	averageSpeed() int
}

type Car struct {
	maxSpeed int
	minSpeed int
}

func (c Car) averageSpeed() int {
	return (c.minSpeed + c.maxSpeed) / 2
}

type Bicycle struct {
	maxSpeed int
	minSpeed int
}

func (b Bicycle) averageSpeed() int {
	return (b.minSpeed + b.maxSpeed) / 2
}

func main() {
	fmt.Println("Go")

	sccot := Bicycle{5, 45}
	vw := Car{10, 120}

	printAverage(vw)
	printAverage(sccot)
}

func printAverage(cd Characteristics) {
	fmt.Println("This thing can go average on :", cd.averageSpeed())
}
