package main

import "fmt"

func handlepanic() {
	fmt.Println("Hello World")
}
func main() {
	defer handlepanic()
	panic("Hello panic")
}
