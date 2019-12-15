package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	var sum map[int]int
	sum[0] = 1
}
