package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "zaky",
		"address": "depok",
	}
	person["age"] = "24"

	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["buku1"] = "isi buku1"
	book["buku2"] = "isi buku2"

	fmt.Println(book)
	delete(book, "buku2")
	fmt.Println(book)
}
