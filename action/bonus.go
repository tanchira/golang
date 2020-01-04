package main

import "fmt"

func main() {
	want = 60000
	Hugesales := 50000 / 100 * 30
	sales := 49999 / 100 * 10
	if Hugesales >= 50000 {
		fmt.Println("ให้โบนัส ", Hugesales, "บาท")
	} else {
		fmt.Println("ให้โบนัส ", sales, "บาท")
	}
}
