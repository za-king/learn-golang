package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("tiga koma lima", 3.5)

	fmt.Println("benar", true)
	fmt.Println(len("hello world"))
	fmt.Println("hello world"[0])

	var fName string

	fName = "zaky S"
	fmt.Println(fName)
	fName = "zaky S2"
	fmt.Println(fName)

	var lastName = "syukur"

	fmt.Println(lastName)

	var int = 36

	fmt.Println(int)

	name := "zaky"

	fmt.Println(name)

	var (
		kata1 = "ini kata 1"
		kata2 = "ini kata 2"
	)

	fmt.Println(kata1, kata2)
}
