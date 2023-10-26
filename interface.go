package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(hasname HasName) {
	fmt.Println("hello ", hasname.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func main() {

	person1 := Person{Name: "andi"}
	sayHello(person1)
}
