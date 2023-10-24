package main

import "fmt"

func main() {

	type noKTP string
	type menikah bool

	var noKTPZaky noKTP = "124131231W"
	fmt.Println(noKTPZaky)

	var statusZaky menikah = false
	fmt.Println(statusZaky)
}
