package main

import "fmt"

func main() {
	salary := 4000
	deduct := 10000 / 100 * 10
	deduct2 := 5001 / 100 * 5
	if deduct >= 10000 {
		fmt.Println(deduct)
	} else if deduct2 >= 10000 {
		fmt.Println(deduct2)
	}
}
