package main

import "fmt"

func main() {
	salary := 10000
	deduct := 10000 / 100 * 10
	deduct2 := 5001 / 100 * 5
	if salary >= 10000 {
		fmt.Println(deduct)
	} else if salary >= 5001 {
		fmt.Println(deduct2)
	}
}
