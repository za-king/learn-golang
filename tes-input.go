package main

import "fmt"

func main() {
	var fName string
	fmt.Println("enter first name :")
	fmt.Scanln(&fName)

	fmt.Println("fullname is " + fName)
}
