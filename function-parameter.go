package main

import "fmt"

func main() {
	sayHelloWorld("zaky", "syukur")
}

func sayHelloWorld(firstName string, lastName string) {

	fmt.Println("hello", firstName, lastName)
}
