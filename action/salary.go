package main

import "fmt"

func main() {
	salary := 10000
	if 10000 > 10000 {
		deduct := 10000 / 100 * 10
		fmt.Println(deduct)

	} else if 5001 >= 10000 {
		deduct2 := 5001 / 100 * 5
		fmt.Println(deduct2)
	} else {
		fmt.Println("not deduct")
	}
}
