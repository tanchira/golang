package main

import "fmt"

func main() {
	cost := make([]int, 5, 10)
	fmt.Println(cost)
	fmt.Println(len(cost))
	fmt.Println(cap(cost))
}
