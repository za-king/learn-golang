package main

import "fmt"

func main() {
	hasil := funcReturnValue("zaky")
	fmt.Println(hasil)
}

func funcReturnValue(name string) string {
	return "hello " + name
}
