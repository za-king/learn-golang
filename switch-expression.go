package main

import "fmt"

func main() {
	name := "zaky"

	switch name {
	case "zaky":
		fmt.Println("ini zaky")
	case "syukur":
		fmt.Println("ini syukur")

	default:
		fmt.Println("ini default")
	}

}
