package main

import "fmt"

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}
	subtract := func(num1, num2 int) int {
		return num1 - num2
	}
	num1 := compute(sum)
	num2 := compute(subtract)
	fmt.Println(num1)
	fmt.Println(num2)
}
