package main

import "fmt"

func main() {
	want := 51896
	Hugesales := want / 100 * 30
	sales := 49999 / 100 * 10
	if Hugesales >= 50000 {
		fmt.Println(Hugesales, "ให้โบนัส ")
	} else if >= 49999{
		fmt.Println(sales, "ให้โบนัส ")
	}
}
