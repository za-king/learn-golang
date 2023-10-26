package main

import "fmt"

type Filter func(name string) string

func getsayHelloWithFilter(name string, filter Filter) {

	nameFiltered := filter(name)

	fmt.Println("hai " + nameFiltered)
}

func spamFiltered(name string) string {
	if name == "anjing" {
		return ",,,"

	} else {
		return name
	}
}

func main() {
	getsayHelloWithFilter("budi", spamFiltered)
}
