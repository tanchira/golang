package main

import "fmt"

func main() {
	number := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(number[5])
	number[5] = 15
	fmt.Println(number[5])
	lenght := len(number)
	fmt.Println("Lenght  =", lenght)
}
