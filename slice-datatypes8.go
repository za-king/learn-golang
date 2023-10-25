package main

import "fmt"

func main() {

	month := [...]string{
		"jan", "feb", "mar", "apr", "mei", "jun", "jul", "agus", "sep", "okt", "nov", "des",
	}

	var slice1 = month[3:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	dayslice1 := days[2:]
	fmt.Println(days)
	fmt.Println(dayslice1)
	dayslice2 := append(dayslice1, "ditambah")
	fmt.Println(dayslice2)
	fmt.Println(dayslice1)

	iniarray := [...]int{1, 2, 3}
	inislice := []int{1, 2, 3}
}
