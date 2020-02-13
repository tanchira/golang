package main

import "fmt"

func main() {
	text := 0
	for {
		fmt.Println(text)
		text = text + 1
		if text >= 3 {
			break
		}
	}

}
