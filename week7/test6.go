package main

import "fmt"

func handlePanic() {
	fmt.Println("Hello World")
}
func main() {
	defer handlePanic()
	panic("Hello panic")
}
