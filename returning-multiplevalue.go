package main

import "fmt"

func getFullName() (string, string) {
	return "zaky", "syukur"
}

func main() {
	firstName, _ := getFullName()

	fmt.Println(firstName)

}
