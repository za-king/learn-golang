package main

import "fmt"

func main() {

	name := "zaky"
	if true == true {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	if length := len(name); length > 3 {
		fmt.Println("besar")
	}
}
