package main

import "fmt"

func main() {

	const a = 10
	const b = 20
	const c = a + b
	const d = b % 2

	fmt.Println(c)
	fmt.Println(d)

	var e = 1

	e++

	fmt.Println(e)
}
