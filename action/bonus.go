package main

import "fmt"

func main() {
	want := 50000
	Hugesales := 50000 / 100 * 30
	if want >= 50001 {
		fmt.Println(Hugesales)
	}
}
