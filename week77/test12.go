package main

import "fmt"

func say() {
	fmt.Println("Hi Goku")
}

func main() {
	defer say()
	fmt.Println("Hello world")
}
