package main

import "fmt"

func main() {
	for text := 1; text <= 20; text = text + 1 {
		if text%2 == 0 {
			fmt.Println(text, "even")
		} else {
			fmt.Println(text, "odd")
		}

	}
}
