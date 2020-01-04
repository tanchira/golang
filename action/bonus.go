package main

import "fmt"

func main() {
	want := 60000
	Hugesales := want / 100 * 30
	sales := 49999 / 100 * 10
	if Hugesales >= 60000 {
		fmt.Println("ให้โบนัส ", Hugesales, "บาท")
	} else if sales >= 49999 {
		fmt.Println("ให้โบนัส ", sales, "บาท")
	}
}
