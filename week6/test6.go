package main

import "fmt"

func writeline(text ...interface{}) {
	for _, v := range text {
		fmt.Println(v)
	}
}

func main() {
	writeline(1, 3.14, "Hello ", true)
}
