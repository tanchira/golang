package main

import "fmt"

func main() {
	number := [5]int{0, 1, 2, 3, 4}
	fmt.Println(number[1])
	number[1] = 10
	fmt.Println(number[1])
	lenght := len(number)
	fmt.Println("Lenght  =", lenght)
}
