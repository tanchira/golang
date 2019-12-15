package main

import "fmt"

func suntract(number int) {
	number = number - 1
	return number
}

func main() {
	sum := 10
	suntract(sum)
	fmt.Println(sum)
}
