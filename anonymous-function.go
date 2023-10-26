package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {

	if blackList(name) {
		fmt.Println("boleh masuk " + name)
	} else {
		fmt.Println("tidak boleh masuk " + name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("budi", blacklist)
}
