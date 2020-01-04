package main

import "fmt"

func main() {
	want := 50000
	Hugesales := want / 100 * 30
	if Hugesales >= 50000 {
		fmt.Println(Hugesales, "ให้โบนัส ")
	}
}
