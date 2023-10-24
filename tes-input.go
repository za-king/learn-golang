package main

import "fmt"

func main() {
	var fName string
	fmt.Println("enter first name :")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("enter last name :")
	fmt.Scanln(&lName)

	fmt.Println("fullname is " + fName + " " + lName)
}
