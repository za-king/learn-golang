package main

import "fmt"

func upps(i int) interface{} {
	if i == 1 {
		return "wow"
	} else if i == 2 {
		return "wew"
	} else {
		return "waduuh"
	}
}

func main() {
	kosong := upps(4)

	fmt.Println(kosong)
}
