package main

import "fmt"

func main() {
	angka := 1

	for angka <= 10 {
		fmt.Println("ini angka ke = ", angka)
		angka++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	slice := []string{"ini 1", "ini 2", "ini 3"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range

	for i, value := range slice {
		fmt.Println("index", i, "value", value)
	}
}
