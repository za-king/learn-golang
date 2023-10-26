package main

import "fmt"

type Admin struct {
	name string
	age  int
}

func (hi Admin) sayHi(from string) {
	fmt.Println("it is", from, "hai my name is", hi.name)
}

func main() {
	admin1 := Admin{
		name: "zaky",
	}

	admin1.sayHi("budi")
}
