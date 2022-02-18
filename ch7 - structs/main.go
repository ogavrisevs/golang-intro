package main

import "fmt"

type Person struct {
	name    string
	surname string
}

type Group struct {
	name    string
	persons [3]Person
}

func (p Person) printNameNice() {
	fmt.Println("-=", p.name, "=-")
}

func main() {
	fmt.Println("Ch7")

	nullPointer()
	simpleObj()
	directObj()
	objMethod()
	embeddedObj()
}

func nullPointer() {
	fmt.Println("nullPointer")

	var p0 Person
	fmt.Println(p0)
	fmt.Println("Person.name : ", p0.name, "Person.surname : ", p0.surname)

	p0.name = "John"
	p0.surname = "Don"
	fmt.Println(p0)
	fmt.Println("Person.name : ", p0.name, "Person.surname : ", p0.surname)
}

func simpleObj() {
	fmt.Println("simpleObj")

	p1 := new(Person) //pointer
	fmt.Println(p1)
	p1.name = "Vasa"
	p1.surname = "Pupkin"

	fmt.Println(p1)
}

func directObj() {
	fmt.Println("directObj")

	p2 := Person{name: "Kuza", surname: "Murza"}
	fmt.Println(&p2)
}

func objMethod() {
	fmt.Println("objMethod")
	p := Person{name: "Vasa", surname: "Pupkin"}
	p.printNameNice()

}

func embeddedObj() {
	fmt.Println("embeddedObj")
	p1 := Person{name: "Kuza", surname: "Murza"}
	p2 := Person{name: "Katja", surname: "Batja"}

	gr1 := Group{name: "first class"}
	gr1.persons[0] = p1
	gr1.persons[1] = p2

	fmt.Println(gr1)
}
