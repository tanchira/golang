package main

import "fmt"

func main() {
	var p1 *int
	fmt.Println("p1 :", p1)
	p2 := new(int)
	fmt.Println("p2 :", p2)
}
