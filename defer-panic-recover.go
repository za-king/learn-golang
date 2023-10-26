package main

import "fmt"

func logging() {
	fmt.Println("ini logging")
}

func endApp() {

	message := recover()
	if message != nil {
		fmt.Println("mesage error :", message)
	}
	fmt.Println("app selesai")
}

func runApplication() {
	defer logging()
	fmt.Println("run app")
}

func runPanic(error bool) {
	defer endApp()
	if error {
		panic("ada yang error")
	}

	fmt.Println("app berjalan")

}

func main() {
	runApplication()

	runPanic(false)
}
