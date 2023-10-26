package main

import "fmt"

type Customer struct {
	name, addres string
	age          int
}

func main() {

	var cus1 Customer

	cus1.name = "zaky"
	cus1.addres = "deppok"
	cus1.age = 24

	cus2 := Customer{
		name:   "zaky",
		addres: "depok",
		age:    24,
	}

	cus3 := Customer{"zaky3", "depok", 24}

	fmt.Println(cus1)
	fmt.Println(cus2)
	fmt.Println(cus3)
}
