package main

import "fmt"

func main() {
	for essay := 1; essay <= 20; essay = essay + 1 {
		if essay%2 == 0 {
			fmt.Println(essay, "even")
		} else {
			fmt.Println(essay, "odd")
		}

	}
}
