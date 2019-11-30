package main

import "fmt"

func main() {
	n, e := fmt.Printf("Hello", "World", 123, 456, "\n")
	fmt.Printf("number of bytes written :", n, "\n")
	fmt.Printf("write error encountered :", e, "\n")
}
