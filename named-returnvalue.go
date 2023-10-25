package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {

	firstName = "muhammad"
	middleName = "zaky"
	lastName = "syukur"

	return
}

func main() {
	a, b, c := getFullName2()

	fmt.Println(a, b, c)
}
