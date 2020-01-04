package main

import "fmt"

func main() {
	want := 50000
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	if want >= 50000 {
		fmt.Println(Hugesales)
	} else {
		fmt.Println(sales)
	}
}
