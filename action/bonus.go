package main

import "fmt"

func main() {
	want := 60000
	Hugesales := want / 100 * 30
	sales := 49999 / 100 * 10
	if Hugesales >= 50001 {
		fmt.Println("ให้โบนัส ", Hugesales, "บาท")
	} else {
		fmt.Println("ให้โบนัส ", sales, "บาท")
	}
}
